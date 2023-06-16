package event

type Event struct {
	E string
}

func (Event) Process(e Event){
	//fmt.Println("Consumer: ",e.E)
}

