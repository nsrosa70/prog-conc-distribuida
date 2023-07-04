package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	HelloServerTCP()

	//HelloServerUDP()

	_, _ = fmt.Scanln()
}

func HelloServerTCP() {

	// define o endpoint do servidor TCP
	r, err := net.ResolveTCPAddr("tcp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// cria um listener TCP
	ln, err := net.ListenTCP("tcp", r)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("Servidor TCP aguardando conexão...")

	// aguarda/aceita conexão
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err.(net.Error))
		os.Exit(0)
	}

	// fecha conexão
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err.(net.Error))
			os.Exit(0)
		}
	}(conn)

	// recebe e processa requests
	for {
		// recebe request terminado com '\n'
		req, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// processa request
		rep := strings.ToUpper(req)

		// envia resposta
		_, err = conn.Write([]byte(rep + "\n"))
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}
}

func HelloServerUDP() {
	req := make([]byte, 1024)
	rep := make([]byte, 1024)

	// define o endpoint do servidor UDP
	addr, err := net.ResolveUDPAddr("udp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// prepara o endpoint UDP para receber requests
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// fecha conn
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(conn)

	fmt.Println("Servidor UDP aguardando requests...")

	for {
		// recebe request
		_, addr, err := conn.ReadFromUDP(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// processa request
		rep = []byte(strings.ToUpper(string(req)))

		// envia reposta
		_, err = conn.WriteTo(rep, addr)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}
}
