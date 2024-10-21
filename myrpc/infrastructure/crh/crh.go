package crh

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
	"test/shared"
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

/*
func (crh *CRH) SendReceive(msgToServer []byte) []byte {
	var err error

	for i := 0; i < shared.MaxConnectionAttempts; i++ {
		crh.Connection, err = net.Dial("tcp", crh.Host+":"+strconv.Itoa(crh.Port))
		if err == nil {
			fmt.Println("CRH:: SendReceive", i)
			break
		}
		time.Sleep(10 * time.Millisecond)
		if i == shared.MaxConnectionAttempts {
			log.Fatal("CRH 0:: Number Max of attempts achieved...")
		}
	}

	// 2: send message's size
	sizeMsgToServer := make([]byte, 4)
	l := uint32(len(msgToServer))
	binary.LittleEndian.PutUint32(sizeMsgToServer, l)
	_, err = crh.Connection.Write(sizeMsgToServer)
	if err != nil {
		log.Fatalf("CRH 1:: %s", err)
	}

	// 3: send message
	_, err = crh.Connection.Write(msgToServer)
	if err != nil {
		log.Fatalf("CRH 2:: %s", err)
	}

	// 4: receive message's size
	sizeMsgFromServer := make([]byte, 4)
	_, err = crh.Connection.Read(sizeMsgFromServer)
	if err != nil {
		log.Fatalf("CRH 3:: %s", err)
	}
	sizeFromServerInt := binary.LittleEndian.Uint32(sizeMsgFromServer)

	//5: receive reply
	msgFromServer := make([]byte, sizeFromServerInt)
	_, err = crh.Connection.Read(msgFromServer)
	if err != nil {
		log.Fatalf("CRH 4:: %s", err)
	}

	crh.Connection.Close()

	return msgFromServer
}
*/

func (crh *CRH) SendReceive(msgToServer []byte) []byte {
	var err error
	s := 0
	msgFromServer := []byte{}

	for {
		switch s {
		case 0: // open connection
			for i := 0; i < shared.MaxConnectionAttempts; i++ {
				crh.Connection, err = net.Dial("tcp", crh.Host+":"+strconv.Itoa(crh.Port))
				if err == nil {
					s = 1
					break
				} else {
					if i == shared.MaxConnectionAttempts {
						log.Fatal("CRH 0:: Number Max of attempts achieved...")
					}
				}
			}
		case 1:
			// 2: send message's size
			sizeMsgToServer := make([]byte, 4)
			l := uint32(len(msgToServer))
			binary.LittleEndian.PutUint32(sizeMsgToServer, l)
			_, err = crh.Connection.Write(sizeMsgToServer)
			if err != nil {
				log.Fatalf("CRH 1:: %s", err)
			}
			s = 2
		case 2:
			// 3: send message
			_, err = crh.Connection.Write(msgToServer)
			if err != nil {
				log.Fatalf("CRH 2:: %s", err)
			}
			s = 3
		case 3:
			// 4: receive message's size
			sizeMsgFromServer := make([]byte, 4)
			_, err = crh.Connection.Read(sizeMsgFromServer)
			if err != nil {
				log.Fatalf("CRH 3:: %s", err)
			}
			sizeFromServerInt := binary.LittleEndian.Uint32(sizeMsgFromServer)

			//5: receive reply
			msgFromServer = make([]byte, sizeFromServerInt)
			_, err = crh.Connection.Read(msgFromServer)
			if err != nil {
				log.Fatalf("CRH 4:: %s", err)
			}
			s = 4
		case 4:
			crh.Connection.Close()
			return msgFromServer
		}
	}
	return msgFromServer
}
