package main

import (
	"fmt"
	"runtime"
)

func main() {

	//cores_virtuais = x * cores_fisicos  ou x = cores_virtuais / cores_fisicos
	//where x = número de hardware threads por core

	// Comando MacOs
	//sysctl hw.physicalcpu hw.logicalcpu

	cores_fisicos := 2 // Minha máquina tem 2 cores
	cpus_logicas := runtime.NumCPU()
	hardware_threads_por_core_fisico := cpus_logicas / cores_fisicos
	//os_threads := cpus_logicas
	cores_virtuais := cpus_logicas
	default_go_max_procs := runtime.GOMAXPROCS(2) // "1" é aleatório, pois esta função retorna o valor anterior
	go_max_procs := 1

	fmt.Println("Cores Físicos            :", cores_fisicos)
	fmt.Println("Cores Virtuais           :", cores_virtuais)
	fmt.Println("CPUs lógicas             :", cpus_logicas)
	fmt.Println("Hardware threads por Core:", hardware_threads_por_core_fisico) //
	fmt.Println("Threads do SO (direct)  :", default_go_max_procs)
	fmt.Println("Threads do SO (desejados):", go_max_procs)
}
