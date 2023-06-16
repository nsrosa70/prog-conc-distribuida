package main

import (
	"fmt"
	"sync"
	"time"
)

func fluxodadireita(ponte sync.Mutex) {
	for {
		fmt.Println("Direita:: Esperando")
		ponte.Lock()
		fmt.Println("Direita:: Entrou na ponte")
		ponte.Unlock()
		fmt.Println("Direita:: Saiu da ponte")
		time.Sleep(3 * time.Second)
	}
}

func fluxodaesquerda(ponte sync.Mutex) {
	for {
		fmt.Println("Esquerda:: Esperando")
		ponte.Lock()
		fmt.Println("Esquerda:: Entrou na ponte")
		ponte.Unlock()
		fmt.Println("Esquerda:: Saiu da ponte")
		time.Sleep(3 * time.Second)
	}
}

func main() {
	ponte := sync.Mutex{}

	go fluxodadireita(ponte)
	go fluxodaesquerda(ponte)

	fmt.Scanln()
}
