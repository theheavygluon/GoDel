package main

import (
	"GoDel"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

func main() {
	var qc = GoDel.QuantumCircuit{Qubits: 2}
	qc.ApplyGate(GoDel.H, 0)
	qc.ApplyGate(GoDel.Z, 1)
	qc.ApplyGate(GoDel.X, 1)
	fmt.Println(GoDel.H.MulElem(1/math.Sqrt(2)))
	fmt.Println(GoDel.H.Caps()) // get domension of matrix
	fmt.Println(GoDel.H.ColView(0)) // view first col of H
	//TODO
	// find a way to make the matrix construction short
	var precustomMatrixData = make([]float64, 6)
	var customMatrixData = [6]float64{1, 1, 1, 1, 1, 1}
	for i := range customMatrixData {
		precustomMatrixData[i] = customMatrixData[i]
	}
	// END HERE
	customMatrix := mat.NewDense(3, 2, precustomMatrixData)
	qc.ApplyGate(customMatrix, 1) // add custom matrix to the gates
	fmt.Println(qc.Gates)
}
