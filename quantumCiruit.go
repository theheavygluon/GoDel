package GoDel

import (
	"errors"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
)

type Gate struct {
	qubit  int
	matrix mat.Matrix
}
type QuantumCircuit struct {
	Gates []Gate
	Qubits int
}

// TODO matrix H
// find a way to make matrix contrauctions short
var preHData = make([]float64, 4)
var hData = [4]float64{1, 1, 1, -1}

var H = mat.NewDense(2, 2, preHData)
// end of h

// TODO matrix X
// find a way to make matrix contrauctions short

var preXData = make([]float64, 4)
var xData = [4]float64{0, 1, 1, 0}

var X = mat.NewDense(2, 2, preXData)

// end of x

// TODO matrix Z
// find a way to make matrix contrauctions short
var preZData = make([]float64, 4)
var zData = [4]float64{1, 0, 0, -1}

var Z = mat.NewDense(2, 2, preZData)

// end of z

func init() {
	// init method to set up everything in the package
	// adding real values to H
	for i := range hData {
		preHData[i] = hData[i]
	}
	// adding real values to X
	for i := range xData {
		preXData[i] = xData[i]
	}
	// adding real values to Z
	for i := range zData {
		preZData[i] = zData[i]
	}
}

func (qc *QuantumCircuit) ApplyGate(gate mat.Matrix, qubit int) {
	if qubit >= qc.Qubits{
		err := errors.New("qubit: qubit out of range for  the circuit")
		log.Println(err)
		return
	}
	var newGate = Gate{qubit: qubit, matrix: gate}
	qc.Gates = append(qc.Gates, newGate)
}

func (qc *QuantumCircuit) GetQubits() int{
	fmt.Println(qc)
	return qc.Qubits
}