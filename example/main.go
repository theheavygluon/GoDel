package main

import (
	"GoDel"
	"fmt"
)

func main()  {
	var qc = GoDel.QuantumCircuit{}
	fmt.Println(GoDel.H)
	qc.ApplyGate(GoDel.H, 2)
}
