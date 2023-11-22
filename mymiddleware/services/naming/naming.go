package naming

import (
	"mymiddleware/distribution/clientproxy"
)

type NamingService struct {
	Repository map[string]clientproxy.ClientProxy
}

func (naming *NamingService) Register(name string, proxy clientproxy.ClientProxy) bool {
	r := false

	// check if repository is already created
	if len(naming.Repository) == 0 {
		naming.Repository = make(map[string]clientproxy.ClientProxy)
	}
	// check if the service is already registered
	_, ok := naming.Repository[name]
	if ok {
		r = false // service already registered
	} else { // service not registered
		naming.Repository[name] = clientproxy.ClientProxy{TypeName: proxy.TypeName, Host: proxy.Host, Port: proxy.Port}
		r = true
	}

	return r
}

func (naming NamingService) Lookup(name string) clientproxy.ClientProxy {

	return naming.Repository[name]
}

func (naming NamingService) List(name string) map[string]clientproxy.ClientProxy {

	return naming.Repository
}
