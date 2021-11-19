package GoDel

import (
	"errors"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
	"reflect"
)

var zero = mat.NewVecDense(2, []float64{1, 0})
var one = mat.NewVecDense(2, []float64{0, 1})

type Gate struct {
	qubit  int
	matrix mat.Matrix
}
type QuantumCircuit struct {
	Gates  []Gate
	Qubits int
}

// H create matrix H (haddaman gate)
/*
Wll be scalled with (1 / sqrt(2)) -> (1 / sqrt(2)) * H
*/
var H = mat.NewDense(2, 2, []float64{
	1, 1,
	1, -1,
})

// end of h

// X
/*
create matrix X (basic NOT gate)
*/
var X = mat.NewDense(2, 2, []float64{
	0, 1,
	1, 0,
})

// end of x

// Z create matrix Z
/*
	matrix z
*/

var Z = mat.NewDense(2, 2, []float64{1, 0, 0, -1})

// end of z


var C = []float64{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 0, 1,
	0, 0, 1, 0,
}
// CN create matrix CN (controled not gate)
/*
this will flip target qubit if controlled qubit is set to 1 else no change to target qubit
*/
var CN = mat.NewDense(4, 4, C)
var CX = mat.NewDense(4, 4, C)

// View
//  View, visualize matrix
func View(matrix mat.Matrix) {
	fmt.Println(mat.Formatted(matrix))

}

// end of CN

// CXXUnirary

func CXX(pie float64) mat.Matrix {
	fmt.Println(pie)
	return mat.NewDense(2, 2, []float64{1, 0, 0, -1})
}
func tensor(initialState []mat.Matrix) mat.Matrix{
	x := initialState[0]
	y := initialState[1]
	View(x)
	View(y)
	m := mat.Dense{}
	m.Kronecker(x,y)
	return &m
}
func CXUnitary(numberQubits int, control int, target int) mat.Matrix {
	var left []mat.Matrix
	var right []mat.Matrix
	identity := mat.NewDense(2, 2, []float64{
		1, 0,
		0, 1,
	})
	for i := 0; i < numberQubits; i++ {
		left = append(left, identity)
		right = append(right, identity)
	}

	transposedZero := mat.TransposeVec{Vector: zero}
	transposedOne := mat.TransposeVec{Vector: one}
	leftControl := mat.NewDense(2, 2, nil)
	leftControl.Mul(zero, transposedZero)
	rightControl := mat.NewDense(2, 2, nil)
	rightControl.Mul(one, transposedOne)

	left[control] = leftControl
	right[control] = rightControl
	right[target] = X
	rightTensor := tensor(right)
	leftTensor := tensor(left)
	var m mat.Dense
	m.Add(rightTensor, leftTensor)
	return &m

}

func init() {
	H.Scale(1/math.Sqrt(2), H) // scale H by (1 / sqrt(2))
}

func addQubits(qbt int, gate mat.Matrix) Gate {
	var newGate = Gate{qubit: qbt, matrix: gate}
	return newGate
}

// ApplyGate
/* ApplyGate contructs a gate, accepts(matrix for the gate and also qubit to apply to )
if the qubit specified is out of range the program will panic

if appllying gate to multilple qubits then pass the qubits as SCLICE
example:

ApplyGate(H,[]int{0,1})

will aplly gate H to qubit index 0 and 1
*/
func (qc *QuantumCircuit) ApplyGate(gate mat.Matrix, qubit interface{}) {
	switch reflect.TypeOf(qubit).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(qubit)
		for i := 0; i < s.Len(); i++ {
			qbt := s.Index(i).Interface().(int)
			if qbt >= qc.Qubits || qbt < 0 {
				err := errors.New("Qubits: qubit out of range for  the circuit")
				panic(err)
			}
			newGate := addQubits(qbt, gate)
			qc.Gates = append(qc.Gates, newGate)
		}
	case reflect.Int:
		s := reflect.ValueOf(qubit)
		qbt := s.Interface().(int)
		if qbt >= qc.Qubits || qbt < 0 {
			err := errors.New("Qubits: qubit out of range for  the circuit")
			panic(err)
		}
		newGate := addQubits(qbt, gate)
		qc.Gates = append(qc.Gates, newGate)
	default:
		err := errors.New("Qubits: wrong data type for the qubit")
		panic(err)
	}
}
