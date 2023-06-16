package proxies

import (
	"mymiddleware/aux"
	"mymiddleware/distribution/clientproxy"
	"mymiddleware/distribution/requestor"
	"reflect"
	"shared"
)

type CalculatorProxy struct {
	Proxy clientproxy.ClientProxy
}

func NewCalculatorProxy() CalculatorProxy {
	p := new(CalculatorProxy)

	p.Proxy.TypeName = reflect.TypeOf(CalculatorProxy{}).String()
	p.Proxy.Host = "localhost"
	//p.Proxy.Port = shared.FindNextAvailablePort()  // TODO
	p.Proxy.Port = shared.CALCULATOR_PORT
	return *p
}

func (proxy CalculatorProxy) Add(p1 int, p2 int) int {

	// prepare invocation
	params := make([]interface{},2)
	params[0] = p1
	params[1] = p2
	request := aux.Request{Op:"Add",Params:params}
	inv := aux.Invocation{Host:proxy.Proxy.Host,Port:proxy.Proxy.Port,Request:request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})

	return int(ter[0].(float64))
}

func (proxy CalculatorProxy) Sub(p1 int, p2 int) int {

	// prepare invocation
	params := make([]interface{},2)
	params[0] = p1
	params[1] = p2
	request := aux.Request{Op:"Sub",Params:params}
	inv := aux.Invocation{Host:proxy.Proxy.Host,Port:proxy.Proxy.Port,Request:request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})

	return int(ter[0].(float64))
}

func (proxy CalculatorProxy) Mul(p1 int, p2 int) int {

	// prepare invocation
	params := make([]interface{},2)
	params[0] = p1
	params[1] = p2
	request := aux.Request{Op:"Mul",Params:params}
	inv := aux.Invocation{Host:proxy.Proxy.Host,Port:proxy.Proxy.Port,Request:request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})

	return int(ter[0].(float64))
}

func (proxy CalculatorProxy) Div(p1 int, p2 int) int {

	// prepare invocation
	params := make([]interface{},2)
	params[0] = p1
	params[1] = p2
	request := aux.Request{Op:"Div",Params:params}
	inv := aux.Invocation{Host:proxy.Proxy.Host,Port:proxy.Proxy.Port,Request:request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})

	return int(ter[0].(float64))
}


