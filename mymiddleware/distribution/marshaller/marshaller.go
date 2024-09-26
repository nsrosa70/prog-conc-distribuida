package marshaller

import (
	"encoding/gob"
	"encoding/json"
	"log"
	"test/mymiddleware/distribution/miop"
)

type Marshaller struct{}

func (Marshaller) MarshallerFactory() Marshaller {
	gob.Register(miop.Packet{})

	return Marshaller{}
}

func (Marshaller) Marshall(msg miop.Packet) []byte {

	r, err := json.Marshal(msg)
	//r, err := msgpack.Marshal(msg)
	if err != nil {
		log.Fatalf("Marshaller:: Marshall:: %s", err)
	}

	return r
}

func (Marshaller) Unmarshall(msg []byte) miop.Packet {

	r := miop.Packet{}
	err := json.Unmarshal(msg, &r)
	//err := msgpack.Unmarshal(msg, &r)
	if err != nil {
		log.Fatalf("Marshaller:: Unmarshall:: %s", err)
	}
	return r
}
