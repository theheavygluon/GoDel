package GoDel

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

type Gate struct {
	qubit  int
	matrix mat.Matrix
}
type QuantumCircuit struct {
	gates []Gate
	qubits int
}

// h
var preHData = make([]float64, 4)
var HData = [4]float64{1, 1, 1, -1}

var H = mat.NewDense(2, 2, preHData)

// x

func init() {
	for i := range HData {
		fmt.Println(HData[i])
		preHData[i] = HData[i]
	}
}
func (qc QuantumCircuit) ApplyGate(gate mat.Matrix, qubit int) {
	var newGate = Gate{qubit: qubit, matrix: gate}
	qc.gates = append(qc.gates, newGate)
}
