package main

import (
	"aulas/sistemasadaptativos/shared"
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/sys/windows"
	"time"
)

type Managing struct{}
type Managed struct{}
type SelfAdaptive struct{}

func (Managing) Run(ch chan int) {
	c := shared.Red
	for {
		if WasESCKeyPressed() {
			switch c {
			case Red:
				c = Green
			case Green:
				c = Red
			}
			ch <- c
		}
	}
}

func (Managed) Run(ch chan int) {
	c := Red
	for {
		select {
		case c = <-ch:
		default:
		}
		printColor(c)
	}
}

func (SelfAdaptive) Run() {
	ch := make(chan int)

	go Managing{}.Run(ch)
	go Managed{}.Run(ch)

	fmt.Scanln()
}

func printColor(i int) {
	switch i {
	case shared.Red:
		color.Red("My Red " + shared.Behaviour)
	case shared.Red:
		color.Green("My Green " + shared.Behaviour)
	}
	time.Sleep(1 * time.Second)
}

var user32_dll = windows.NewLazyDLL("user32.dll")
var GetKeyState = user32_dll.NewProc("GetKeyState")

func WasESCKeyPressed() bool {
	r1, _, _ := GetKeyState.Call(27) // Call API to get ESC key state.
	return r1 == 65409               // Code for KEY_UP event of ESC key.
}
