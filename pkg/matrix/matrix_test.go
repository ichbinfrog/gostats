package matrix

import (
	"fmt"
	"math"
	"syscall"
	"testing"

	"github.com/ichbinfrog/statistics/pkg/dist"
)

func init() {
	var rlimit syscall.Rlimit
	syscall.Getrlimit(syscall.RLIMIT_DATA, &rlimit)
	rlimit.Max = rlimit.Max / 2
	rlimit.Cur = rlimit.Cur / 4
	syscall.Setrlimit(syscall.RLIMIT_AS, &rlimit)
}

func BenchmarkMatrix(t *testing.B) {
	var data [][]float64
	for i := 1; i < 5; i++ {
		d := dist.Normal{}
		d.Init(0, 1)

		n := int(math.Pow10(i))
		data = make([][]float64, n)
		for j := 0; j < n; j++ {
			data[j] = make([]float64, n)
			for k := 0; k < n; k++ {
				data[j][k] = d.Generate()
			}
		}
		A := &Matrixf64{
			height:     8,
			width:      8,
			Data:       data,
			Transposed: false,
		}

		B := A
		t.Run(fmt.Sprintf("naive_%d", n), func(t *testing.B) {
			NaiveMult(A, B)
		})

		t.Run(fmt.Sprintf("parallel_row_%d", n), func(t *testing.B) {
			ParallelRowMult(A, B)
		})

		t.Run(fmt.Sprintf("strassen_%d", n), func(t *testing.B) {
			StrassenMult(A, B)
		})
	}

}
