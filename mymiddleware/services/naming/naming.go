package naming

import (
	"test/shared"
)

type NamingService struct {
	Repository map[string]shared.IOR
}

func (n *NamingService) Bind(s string, i shared.IOR) bool {
	r := false

	// check if repository is already created
	if len(n.Repository) == 0 {
		n.Repository = make(map[string]shared.IOR)
	}
	// check if the service is already registered
	_, ok := n.Repository[s]
	if ok {
		r = false // service already registered
	} else { // service not registered
		n.Repository[s] = shared.IOR{TypeName: i.TypeName, Host: i.Host, Port: i.Port}
		r = true
	}

	return r
}

func (n NamingService) Find(s string) shared.IOR {

	return n.Repository[s]
}

func (n NamingService) List() map[string]shared.IOR {

	return n.Repository
}
