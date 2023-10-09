package srh

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

type SRH struct {
	Host       string
	Port       int
	Connection net.Conn
}

var ln net.Listener

//var conn net.Conn
var err error

func NewSRH(h string, p int) *SRH {
	r := new(SRH)

	r.Host = h
	r.Port = p
	r.Connection = nil

	return r
}

func (srh *SRH) Receive() []byte {

	if srh.Connection == nil {
		// 1: create listener
		ln, err = net.Listen("tcp", srh.Host+":"+strconv.Itoa(srh.Port))
		if err != nil {
			log.Fatalf("HERE 1:: SRH:: %s", err)
		}

		// 2: accept connection
		srh.Connection, err = ln.Accept()
		if err != nil {
			log.Fatalf("HERE 2:: SRH:: %s", err)
		}
	}

	// 3: receive message's size
	size := make([]byte, 4)
	_, err = srh.Connection.Read(size)
	if err != nil {
		ne, _ := err.(net.Error)
		if strings.Contains(ne.Error(), "wsarecv") {
			return []byte{}
		} else {
			log.Fatalf("SRH:: %s", err)
		}
	}
	sizeInt := binary.LittleEndian.Uint32(size)

	// 4: receive message
	msg := make([]byte, sizeInt)
	_, err = srh.Connection.Read(msg)
	ne, _ := err.(net.Error)
	if strings.Contains(ne.Error(), "wsarecv") {
		return []byte{}
	} else {
		log.Fatalf("SRH:: %s", err)
	}
	return msg
}

func (srh *SRH) Send(msgToClient []byte) {

	if srh.Connection == nil {
		fmt.Println("SRH:: Connection not opened")
		os.Exit(0)
	}

	// 1: send message's size
	size := make([]byte, 4)
	l := uint32(len(msgToClient))
	binary.LittleEndian.PutUint32(size, l)
	_, err = srh.Connection.Write(size)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	// 2: send message
	_, err = srh.Connection.Write(msgToClient)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	// 3: close connection
	//srh.Connection.Close()
	//ln.Close()
}

func handler(c net.Conn) {

}
