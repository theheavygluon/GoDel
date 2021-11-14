package main

import (
	"GoDel"
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	var qc = GoDel.QuantumCircuit{Qubits: 2}
	qc.ApplyGate(GoDel.H, 1)

	//TODO
	// find a way to make the matrix construction short
	var precustomMatrixData = make([]float64, 6)
	var customMatrixData = [6]float64{1, 1, 1, 1, 1, 1}
	for i := range customMatrixData {
		precustomMatrixData[i] = customMatrixData[i]
	}
	// END HERE
	customMatrix := mat.NewDense(3, 2, precustomMatrixData)
	qc.ApplyGate(customMatrix, 2) // add custom matrix to the gates
	fmt.Println(qc.Gates)
}
