package srh

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
)

type SRH struct {
	Host       string
	Port       int
	Connection net.Conn
	Ln         net.Listener
}

// var conn net.Conn
var err error

func NewSRH(h string, p int) *SRH {
	r := new(SRH)

	r.Host = h
	r.Port = p
	r.Connection = nil
	// 1: create listener & accept connection
	r.Ln, err = net.Listen("tcp", h+":"+strconv.Itoa(p))
	if err != nil {
		log.Fatalf("SRH 0:: %s", err)
	}

	return r
}

/*
func (srh *SRH) Receive() []byte {

	srh.Connection, err = srh.Ln.Accept()
	if err != nil {
		log.Fatalf("SRH 1:: %s", err)
	}

	// 2: receive message's size
	size := make([]byte, 4)
	_, err = srh.Connection.Read(size)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			srh.Connection.Close()
			return nil
		} else {
			log.Fatalf("SRH 2:: %s", err)
		}
	}
	sizeInt := binary.LittleEndian.Uint32(size)

	// 3: receive message
	msg := make([]byte, sizeInt)
	_, err = srh.Connection.Read(msg)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			srh.Connection.Close()
			return nil
		} else {
			log.Fatalf("SRH 3:: %s", err)
		}
	}
	return msg
}

func (srh *SRH) Send(msgToClient []byte) {

	// 2: send message's size
	size := make([]byte, 4)
	l := uint32(len(msgToClient))
	binary.LittleEndian.PutUint32(size, l)
	_, err = srh.Connection.Write(size)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			srh.Connection.Close()
			return
		} else {
			log.Fatalf("SRH 4:: %s", err)
		}
	}

	// 3: send message
	_, err = srh.Connection.Write(msgToClient)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			srh.Connection.Close()
			srh.Ln.Close()
			return
		} else {
			log.Fatalf("SRH 5:: %s", err)
		}
	}
	//defer srh.Connection.Close()
	//defer srh.Ln.Close()
}
*/

func (srh *SRH) Receive() []byte {

	srh.Connection, err = srh.Ln.Accept()
	if err != nil {
		log.Fatalf("SRH 1:: %s", err)
	}

	// 2: receive message's size
	size := make([]byte, 4)
	_, err = srh.Connection.Read(size)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			srh.Connection.Close()
			return nil
		} else {
			log.Fatalf("SRH 2:: %s", err)
		}
	}
	sizeInt := binary.LittleEndian.Uint32(size)

	// 3: receive message
	msg := make([]byte, sizeInt)
	_, err = srh.Connection.Read(msg)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			srh.Connection.Close()
			return nil
		} else {
			log.Fatalf("SRH 3:: %s", err)
		}
	}
	return msg
}

func (srh *SRH) Send(msgToClient []byte) {

	// 2: send message's size
	size := make([]byte, 4)
	l := uint32(len(msgToClient))
	binary.LittleEndian.PutUint32(size, l)
	_, err = srh.Connection.Write(size)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			srh.Connection.Close()
			return
		} else {
			log.Fatalf("SRH 4:: %s", err)
		}
	}

	// 3: send message
	_, err = srh.Connection.Write(msgToClient)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			srh.Connection.Close()
			srh.Ln.Close()
			return
		} else {
			log.Fatalf("SRH 5:: %s", err)
		}
	}
	//defer srh.Connection.Close()
	//defer srh.Ln.Close()
}
