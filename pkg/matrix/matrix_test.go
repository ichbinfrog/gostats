package matrix

import (
	"log"
	"testing"
)

func TestMatrix(t *testing.T) {
	m := &Matrix{
		height:     3,
		width:      3,
		Data:       [][]float64{[]float64{1, 2, 3}, []float64{1, 2, 3}, []float64{1, 2, 3}},
		Transposed: false,
	}
	n := m
	log.Println(m)
	Add(m, n, true)
	log.Println(m)
	Mult(m, n, true)
	log.Println(m)
}
