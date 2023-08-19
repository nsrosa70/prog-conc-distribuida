package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	HelloClientTCP(10)
	//HelloClientUDP(10)
}

func HelloClientTCP(n int) {

	// retorna o endereço do endpoint TCP
	r, err := net.ResolveTCPAddr("tcp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// connecta ao servidor (sem definir uma porta local)
	conn, err := net.DialTCP("tcp", nil, r)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// fecha connexão
	defer func(conn *net.TCPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(conn)

	for i := 0; i < n; i++ {

		// cria request
		req := "Mensagem " + strconv.Itoa(i)

		// envia mensagem para o servidor
		_, err := fmt.Fprintf(conn, req+"\n")
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// recebe resposta do servidor
		rep, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		fmt.Print(req, " ", rep)
	}
}

func HelloClientUDP(n int) {
	req := make([]byte, 1024)
	rep := make([]byte, 1024)

	// retorna o endereço do endpoint UDP
	addr, err := net.ResolveUDPAddr("udp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// conecta ao servidor -- não cria uma conexão
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// desconecta do servidor
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(conn)

	for i := 0; i < n; i++ {
		// cria request
		req = []byte("Mensagem " + strconv.Itoa(i))

		// envia request ao servidor
		_, err = conn.Write(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// recebe resposta do servidor
		_, _, err := conn.ReadFromUDP(rep)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		fmt.Println(string(req), " -> ", string(rep))
	}
}
