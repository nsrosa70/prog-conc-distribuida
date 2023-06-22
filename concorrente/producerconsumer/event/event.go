package event

import "fmt"

type Event struct {
	E string
}

func (Event) Process(e Event) {
	fmt.Println(e.E)
}
