package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/v3/process"
	"math"
	"os"
	"runtime"
	"time"
)

const MonitorTime = 1000 // milliseconds
const LimitePrime = 10000
const Goal = 50.0 // cpu usage

const BasicNumberOfGoroutines = 1

// const ProcessName = "___go_build_aulas_sistemasadaptativos_ex04.exe"
const ProcessName = "___229go_build_main_go.exe"
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

	// verify controller ouput
	n := runtime.NumGoroutine() - BasicNumberOfGoroutines
	if int(r) > n {
		r = float64(n)
	}
	if int(r) < 0 {
		r = c.Min
	}
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

	// define limits of adaptation
	n := runtime.NumGoroutine() - BasicNumberOfGoroutines
	if int(r) > n {
		r = float64(n)
	}
	return r
}

func Job(ch chan bool, id int) {
	for {
		select {
		case <-ch:
			return // end job
		default:
			findPrimes(LimitePrime) // Do something
		}
	}
}

func CreateJob(ch chan bool, id int) {
	go Job(ch, id)
}

func DeleteJob(ch chan bool) {
	ch <- true
}

func ManageJobs(ch chan bool, cOut int) {
	if cOut > 0 { // increment jobs
		for i := 0; i < cOut; i++ {
			CreateJob(ch, i)
		}
	}
	if cOut < 0 { // decrement jobs
		for i := 0; i < int(math.Abs(float64(cOut))); i++ {
			DeleteJob(ch)
		}
	}
}

func main() {
	ch := make(chan bool)

	CreateJob(ch, 1) // at least an initial job
	c := InitialiseOnOff(-1, 1)
	//c := InitialisePID(-1, 1, 1.0, 0.0125, 0.0125)
	for i := 0; i < SampleSize; i++ {
		n := GetCPUUsageByName(ProcessName)
		cOut := int(c.UpdateOnOff(Goal, n))
		//cOut := int(c.UpdatePID(Goal, n))
		fmt.Printf("%.2f;%.2f;%d; %d\n", Goal, n, runtime.NumGoroutine()-BasicNumberOfGoroutines, cOut)
		//fmt.Printf("%.2f;%.2f;%d\n", Goal, n, runtime.NumGoroutine()-BasicNumberOfGoroutines)
		ManageJobs(ch, cOut)
		time.Sleep(MonitorTime * time.Millisecond)
	}
	//wg.Wait()
	//fmt.Scanln()
}

func findPrimes(limit int) {
	count := 0
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			count++
		}
	}
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

func GetCPUUsage() float64 {
	var r []float64
	r, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return r[0]
}

func GetCPUUsageByName(n string) float64 {
	var r float64

	// Specify the process name you want to monitor
	processName := n

	// Find the process by name
	processes, err := process.Processes()
	if err != nil {
		fmt.Printf("Error getting processes: %v\n", err)
		return r
	}

	var targetProcess *process.Process
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

	r = cpuUsagePercent
	return r
}
