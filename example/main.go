package main

import (
	"GoDel"
	"gonum.org/v1/gonum/mat"
)

func main() {
	var qc = GoDel.QuantumCircuit{Qubits: 2}
	qc.ApplyGate(GoDel.H, []int{1,0}) // apply gate to multiple qubits (dont use int32 or int64, etc...) must use int datatype
	qc.ApplyGate(GoDel.Z, 1)
	qc.ApplyGate(GoDel.X, 1)
	qc.ApplyGate(GoDel.CN, 1)
	qc.ApplyGate(GoDel.CX, 1)
	qc.ApplyGate(GoDel.CXX(120), 1) // using matrix with other params
	x := GoDel.CXUnitary(2,1,1) // using matrix with other params
	GoDel.View(x)

	// customMatrix
	//customMatrix create custom matrix and add to gates
	customMatrix := mat.NewDense(3, 2, []float64{
		1, 1,
		1, 1,
		1, 1,
	})
	qc.ApplyGate(customMatrix, 1) // add custom matrix to the gates
	// end here

}
