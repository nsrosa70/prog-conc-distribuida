package crh

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
	"test/shared"
	"time"
)

type CRH struct {
	Host       string
	Port       int
	Connection net.Conn
}

func NewCRH(h string, p int) *CRH {
	r := new(CRH)
	r.Host = h
	r.Port = p
	r.Connection = nil

	return r
}

func (crh *CRH) SendReceive(msgToServer []byte) []byte {
	var err error

	// 1: connect to servidor
	h := 20 // time between attempts
	if crh.Connection == nil {
		for i := 0; i <= shared.MaxConnectionAttempts; i++ {
			crh.Connection, err = net.Dial("tcp", crh.Host+":"+strconv.Itoa(crh.Port))
			if err == nil {
				break
			} else {
				time.Sleep(time.Duration(h) * time.Millisecond)
				h *= 2
			}
			if i == shared.MaxConnectionAttempts {
				log.Fatal("Number of connection attempts to server exceeded...")
			}
		}
	}

	// 2: send message's size
	sizeMsgToServer := make([]byte, 4)
	l := uint32(len(msgToServer))
	binary.LittleEndian.PutUint32(sizeMsgToServer, l)
	_, err = crh.Connection.Write(sizeMsgToServer)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	// 3: send message
	_, err = crh.Connection.Write(msgToServer)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	// 4: receive message's size
	sizeMsgFromServer := make([]byte, 4)
	_, err = crh.Connection.Read(sizeMsgFromServer)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}
	sizeFromServerInt := binary.LittleEndian.Uint32(sizeMsgFromServer)

	//5: receive reply
	msgFromServer := make([]byte, sizeFromServerInt)
	_, err = crh.Connection.Read(msgFromServer)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	crh.Connection.Close()
	crh.Connection = nil

	return msgFromServer
}
