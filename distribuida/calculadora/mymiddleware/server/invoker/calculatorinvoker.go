package invoker

import (
	"distribuida/calculadora/impl"
	"mymiddleware/distribution/marshaller"
	"mymiddleware/distribution/miop"
	"mymiddleware/infrastructure/srh"
	"shared"
)

type CalculatorInvoker struct {}

func NewCalculatorInvoker() CalculatorInvoker {
	p := new(CalculatorInvoker)

	return *p
}

func (CalculatorInvoker) Invoke(){
	srhImpl := srh.SRH{ServerHost:"localhost",ServerPort:shared.CALCULATOR_PORT}
	marshallerImpl := marshaller.Marshaller{}
	miopPacketReply := miop.Packet{}
	replParams := make([]interface{},1)

	calculatorImpl := impl.Calculadora{}

	for {
		// receive data
		rcvMsgBytes := srhImpl.Receive()

		// unmarshall
		miopPacketRequest := marshallerImpl.Unmarshall(rcvMsgBytes)
		operation := miopPacketRequest.Bd.ReqHeader.Operation

		// demux request
		switch operation {
		case "Add" :
			_p1 := int(miopPacketRequest.Bd.ReqBody.Body[0].(float64))
			_p2 := int(miopPacketRequest.Bd.ReqBody.Body[1].(float64))
			replParams[0] = calculatorImpl.Add(_p1,_p2)
		case "Sub" :
			_p1 := int(miopPacketRequest.Bd.ReqBody.Body[0].(float64))
			_p2 := int(miopPacketRequest.Bd.ReqBody.Body[1].(float64))
			replParams[0] = calculatorImpl.Sub(_p1,_p2)
		case "Mul" :
			_p1 := int(miopPacketRequest.Bd.ReqBody.Body[0].(float64))
			_p2 := int(miopPacketRequest.Bd.ReqBody.Body[1].(float64))
			replParams[0] = calculatorImpl.Mul(_p1,_p2)
		case "Div" :
			_p1 := int(miopPacketRequest.Bd.ReqBody.Body[0].(float64))
			_p2 := int(miopPacketRequest.Bd.ReqBody.Body[1].(float64))
			replParams[0] = calculatorImpl.Div(_p1,_p2)

		}

		// assembly packet
		repHeader := miop.ReplyHeader{Context:"", RequestId: miopPacketRequest.Bd.ReqHeader.RequestId,Status:1}
		repBody := miop.ReplyBody{OperationResult:replParams}
		header := miop.Header{Magic:"MIOP",Version:"1.0",ByteOrder:true,MessageType:shared.MIOP_REQUEST}
		body := miop.Body{RepHeader:repHeader,RepBody:repBody}
		miopPacketReply = miop.Packet{Hdr:header,Bd:body}

		// marshall reply
		msgToClientBytes := marshallerImpl.Marshall(miopPacketReply)

		// send reply
		srhImpl.Send(msgToClientBytes)
	}
}





