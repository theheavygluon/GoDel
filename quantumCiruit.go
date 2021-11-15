package GoDel

import (
	"errors"
	"fmt"
	"gonum.org/v1/gonum/mat"
)

type Gate struct {
	qubit  int
	matrix mat.Matrix
}
type QuantumCircuit struct {
	Gates  []Gate
	Qubits int
}

// H create matrix H
var H = mat.NewDense(2, 2, []float64{1, 1, 1, -1})

// end of h

// X create matrix X
var X = mat.NewDense(2, 2, []float64{0, 1, 1, 0})

// end of x

// Z create matrix Z
var Z = mat.NewDense(2, 2, []float64{1, 0, 0, -1})

// end of z

func init() {

}

func (qc *QuantumCircuit) ApplyGate(gate mat.Matrix, qubit int) {
	if qubit >= qc.Qubits {
		err := errors.New("Qubits: qubit out of range for  the circuit")
		panic(err)
	}
	var newGate = Gate{qubit: qubit, matrix: gate}
	qc.Gates = append(qc.Gates, newGate)
}

func (qc *QuantumCircuit) GetQubits() int {
	fmt.Println(qc)
	return qc.Qubits
}
