package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) GetAge() int {
	return p.Age
}

func main() {
	p := Person{Name: "Jose", Age: 87}
	v := reflect.ValueOf(p)

	op := v.MethodByName("GetAge")
	fmt.Println("Idade: ", p.GetAge())
	fmt.Println("Idade: ", op.Call(nil)[0])
}
