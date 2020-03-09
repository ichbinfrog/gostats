package array

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func populate(a *Arrayf64, n int) {
	for i := 0; i <= n; i++ {
		a.Insert(rand.Float64())
	}
}

func TestArray(b *testing.T) {
	a := Arrayf64{}
	a.Init(4)
	populate(&a, 1000)
	log.Printf("%+v\n", a.Summary())
	a.Center(true)
	log.Printf("%+v\n", a.Summary())
	a.Reduce(true)
	log.Printf("%+v\n", a.Summary())
}

func BenchmarkArraySummary(b *testing.B) {
	for i := 1; i < 8; i++ {
		a := Arrayf64{}
		a.Init(2)

		populate(&a, int(math.Pow(10, float64(i))))
		b.ResetTimer()
		b.Run(fmt.Sprintf("BenchmarkArraySummary_10^%d", i), func(b *testing.B) {
			a.Summary()
		})
	}
}

func BenchmarkArrayApply(b *testing.B) {
	functions := map[string]func(float64) float64{
		"linear": func(v float64) float64 { return v + 1 },
		"exp":    func(v float64) float64 { return math.Exp2(v) },
		"log":    func(v float64) float64 { return math.Log10(v) },
	}

	for name, funct := range functions {
		for i := 1; i <= 5; i++ {
			a := Arrayf64{}
			a.Init(2)

			populate(&a, int(math.Pow(10, float64(i))))
			b.ResetTimer()
			b.Run(fmt.Sprintf("BenchmarkArrayApply_%s_10^%d", name, i), func(b *testing.B) {
				a.apply(funct, true)
			})
		}
	}
}
