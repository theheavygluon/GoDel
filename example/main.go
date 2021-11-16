package main

import (
	"GoDel"
	"gonum.org/v1/gonum/mat"
)

func main() {
	var qc = GoDel.QuantumCircuit{Qubits: 2}
	qc.ApplyGate(GoDel.H, 0)
	qc.ApplyGate(GoDel.Z, 1)
	qc.ApplyGate(GoDel.X, 1)
	qc.ApplyGate(GoDel.CN, 1)
	qc.ApplyGate(GoDel.CXX(120), 1) // using matrix with other params

	//create custom matrix and add to gates
	customMatrix := mat.NewDense(3, 2, []float64{
		1, 1,
		1, 1,
		1, 1,
	})
	qc.ApplyGate(customMatrix, 1) // add custom matrix to the gates
	// end here
}
