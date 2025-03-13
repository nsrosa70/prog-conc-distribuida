package main

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"net"
	"os"
	"strconv"
	"time"
)

const MonitorTime = 1000 // milliseconds
const LimitePrime = 10000
const Goal = 50.0 // cpu usage

// const ProcessName = "___go_build_aulas_sistemasadaptativos_ex04.exe"
const ProcessName = "___11go_build_servidor_go.exe"
const SampleSize = 100

type OnOff struct {
	Max float64
	Min float64
}

type PID struct {
	Kp            float64
	Ki            float64
	Kd            float64
	Max           float64
	Min           float64
	SumError      float64
	Previouserror float64
}

type ManagedSystem struct {
	NClients int
}

type ManagingSystem struct{}

func (m ManagedSystem) GetNClients() int {
	return m.NClients
}

func InitialisePID(min, max, kp, ki, kd float64) PID {
	r := PID{Max: max, Min: min, Kp: kp, Ki: ki, Kd: kd}
	return r
}

func InitialiseOnOff(min, max float64) OnOff {
	r := OnOff{Max: max, Min: min}
	return r
}

func (c *PID) UpdatePID(g, m float64) float64 {
	var r float64
	err := g - m

	p := c.Kp * err
	i := c.Ki * (c.SumError + err)
	d := c.Kd * (c.Previouserror - err)
	c.Previouserror = err
	c.SumError += err

	r = p + i + d

	return r
}

func (c OnOff) UpdateOnOff(g, m float64) float64 {
	var r float64
	err := g - m
	if err > 0 {
		r = c.Max
	} else {
		r = c.Min
	}

	return r
}

func (m *ManagedSystem) Run(chManaging chan int) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	err = listener.(*net.TCPListener).SetDeadline(time.Now().Add(30 * time.Hour)) // TODO
	if err != nil {
		fmt.Println("Error setting deadline:", err)
		return
	}
	defer listener.Close()
	m.NClients = 0
	chHandler := make(chan bool)
	nextMaxClients := 100

	for {
		select {
		case <-chHandler:
			m.NClients--
		case nextMaxClients = <-chManaging:
		default:
			if m.NClients < nextMaxClients {
				conn, err := listener.Accept() // Accept client connection
				if err != nil {
					//fmt.Println("Connection error:", err)
					continue
				}
				go handleClient(conn, chHandler) // Handle client in a new goroutine
				m.NClients++
			}
		}
	}
}

func (ManagingSystem) Run(chManaged chan int, managed *ManagedSystem) {
	chCPUMonitor := make(chan float64)
	//c := InitialiseOnOff(-1, 1)
	c := InitialisePID(-1, 1, 0.1250, 0.0, 0.0)
	//c := InitialisePID(-1, 1, 1.0, 0.5, 1.5)
	go cpuMonitor(chCPUMonitor, ProcessName)

	for i := 0; i < SampleSize; i++ {
		cpuUsage := <-chCPUMonitor
		nClients := managed.GetNClients() // receive #clients
		//cOut := int(c.UpdateOnOff(Goal, n))
		cOut := int(c.UpdatePID(Goal, cpuUsage))
		fmt.Printf("%.2f;%.2f;%v;%v\n", Goal, cpuUsage, nClients, cOut)
		//, cOut)
		chManaged <- cOut
	}
}

func cpuMonitor(chToManaging chan float64, processName string) {
	var targetProcess *process.Process

	processes, err := process.Processes()
	if err != nil {
		fmt.Printf("Error getting processes: %v\n", err)
		return
	}

	for {
		for _, p := range processes {
			name, err := p.Name()
			if err != nil {
				continue
			}
			if name == processName {
				targetProcess = p
				break
			}
		}

		if targetProcess == nil {
			fmt.Printf("Process with name '%s' not found\n", processName)
			os.Exit(0)
		}

		// Measure CPU usage over a period of time
		initialCPU, err := targetProcess.Times()
		if err != nil {
			fmt.Printf("Error getting initial CPU times: %v\n", err)
			os.Exit(0)
		}

		time.Sleep(1 * time.Second) // Wait for 1 second

		finalCPU, err := targetProcess.Times()
		if err != nil {
			fmt.Printf("Error getting final CPU times: %v\n", err)
			os.Exit(0)
		}

		// Calculate CPU usage percentage
		cpuUsage := (finalCPU.User + finalCPU.System) - (initialCPU.User + initialCPU.System)
		cpuUsagePercent := (cpuUsage / 1.0 * 100) / 6
		chToManaging <- cpuUsagePercent
		time.Sleep(MonitorTime * time.Millisecond)
	}
}

func handleClient(conn net.Conn, chManaged chan bool) {
	defer conn.Close()
	//fmt.Println("Client connected:", conn.RemoteAddr())

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			//fmt.Println("Client disconnected:", conn.RemoteAddr())
			chManaged <- true // inform that the client finished
			return
		}
		n, _ := strconv.Atoi(message[:len(message)-1])
		conn.Write([]byte(strconv.Itoa(findPrimes(n)) + "\n"))
	}
}

func findPrimes(limit int) int {
	r := 0
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			r++
		}
	}
	return r
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	managed := ManagedSystem{}
	managing := ManagingSystem{}

	go managed.Run(ch)
	go managing.Run(ch, &managed)

	fmt.Scanln()
}
