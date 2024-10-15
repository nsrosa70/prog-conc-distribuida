package miop

import (
	"test/shared"
)

type Packet struct {
	Hdr Header
	Bd  Body
}

type Header struct {
	Magic       string
	Version     string
	ByteOrder   bool
	MessageType int
	Size        int
}

type Body struct {
	ReqHeader RequestHeader
	ReqBody   RequestBody
	RepHeader ReplyHeader
	RepBody   ReplyBody
}

type RequestHeader struct {
	Context          string
	RequestId        int
	ResponseExpected bool
	ObjectKey        int
	Operation        string
}

type RequestBody struct {
	Body []interface{}
}

type ReplyHeader struct {
	Context   string
	RequestId int
	Status    int
}

type ReplyBody struct {
	OperationResult []interface{}
}

func CreateRequestMIOP(op string, p []interface{}) Packet {
	r := Packet{}

	miopHeader := Header{}
	miopBody := Body{}
	reqHeader := RequestHeader{Operation: op}
	reqBody := RequestBody{Body: p}
	miopBody = Body{ReqHeader: reqHeader, ReqBody: reqBody}

	r.Hdr = miopHeader
	r.Bd = miopBody

	return r
}

func CreateReplyMIOP(params []interface{}) Packet {
	r := Packet{}

	miopHeader := Header{}
	miopBody := Body{}
	repHeader := ReplyHeader{"", 1313, 1} // TODO
	repBody := ReplyBody{OperationResult: params}
	miopBody = Body{RepHeader: repHeader, RepBody: repBody}

	r.Hdr = miopHeader
	r.Bd = miopBody

	return r
}

func ExtractRequest(m Packet) shared.Request {
	i := shared.Request{}

	i.Op = m.Bd.ReqHeader.Operation
	i.Params = m.Bd.ReqBody.Body

	return i
}

func ExtractReply(m Packet) shared.Reply {
	var r shared.Reply

	r.Result = m.Bd.RepBody.OperationResult

	return r
}
