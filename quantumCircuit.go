package GoDel

import (
	"errors"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
	"reflect"
)

type Gate struct {
	qubit  interface{}
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

// CN create matrix CN (controled not gate)
/*
this will flip target qubit if controlled qubit is set to 1 else no change to target qubit
*/
var CN = mat.NewDense(4, 4, []float64{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 0, 1,
	0, 0, 1, 0,
})

// end of CN

// CXX
/*craete a matrix CXX that also accepts inputs
 */
func CXX(pie float64) mat.Matrix {
	fmt.Println(pie)
	return mat.NewDense(2, 2, []float64{1, 0, 0, -1})
}

func init() {
	H.Scale(1/math.Sqrt(2), H) // scale H by (1 / sqrt(2))
}

//ApplyGate
/*
 applyGate contructs a gate, accepts(matrix for the gate and also qubit to apply to )
if the qubit specified is our of range the program will panic

if appllying gate to multilple qubits then pass the qubits as array eample: `[]int32{0,1}` will aplly gate to qubit index 0 and 1
*/

func (qc *QuantumCircuit) ApplyGate(gate mat.Matrix, qubit interface{}) {
	switch reflect.TypeOf(qubit).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(qubit)

		for i := 0; i < s.Len(); i++ {
			y := s.Index(i).Interface().(int)
			if y >= qc.Qubits {
				err := errors.New("Qubits: qubit out of range for  the circuit")
				panic(err)
			}
			var newGate = Gate{qubit: s.Index(i), matrix: gate}
			qc.Gates = append(qc.Gates, newGate)
		}
	case reflect.Int:
		s := reflect.ValueOf(qubit)
		if s.Interface().(int) >= qc.Qubits {
			err := errors.New("Qubits: qubit out of range for  the circuit")
			panic(err)
		}
		var newGate = Gate{qubit: s, matrix: gate}
		qc.Gates = append(qc.Gates, newGate)
	}
}
