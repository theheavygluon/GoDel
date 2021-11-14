package main

import (
	"GoDel"
	"fmt"
)

func main()  {
	var qc = GoDel.QuantumCircuit{}
	qc.ApplyGate(GoDel.H, 2)
	fmt.Println(qc)
}
