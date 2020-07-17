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
	for i := 0; i < n; i++ {
		a.Insert(rand.Float64() * 10)
	}
}

func TestSwilk(b *testing.T) {
	a := Arrayf64{}
	a.Init(Optionf64{
		Degree: 4,
	})
	populate(&a, 4000)
	W := a.ShapiroWilkStatistic()
	log.Println(W, ShapiroWilkSignificance(a.Length, W))
}

func TestSigmoid(b *testing.T) {
	a := Arrayf64{}
	a.Init(Optionf64{
		Degree: 4,
	})
	populate(&a, 100)
	log.Println(a.Data)
	na := a.Sigmoid(false)
	log.Println(na)
}

func TestAggregate(b *testing.T) {
	a := Arrayf64{}
	a.Init(Optionf64{
		Degree:    4,
		Harmonic:  true,
		Geometric: true,
	})
	a.InsertSlice([]float64{1, 2, 5, 3, 6, 78, 35, 2, 1, 1, 2, 3, 2, 1, 2, 2, 2, 2, 3, 5, 4, 4, 5})
	log.Println(a.Data)
	log.Println(a.Aggregate)

	log.Println(a.HarmonicMean(false))
	a.apply(func(v float64) float64 { return v + 10 }, true)
	log.Println(a.Data)

	log.Println(a.HarmonicMean(false))
}

func TestMode(b *testing.T) {
	a := Arrayf64{}
	a.Init(Optionf64{
		Degree: 2,
	})
	a.InsertSlice([]float64{1, 1, 1, 1, 2, 2, 2, 2, 2.2, 3, 2, 2, 2, 3, 5})
	log.Println(a.Mode(false))
	a.apply(func(v float64) float64 { return v + 1 }, true)
	log.Println(a.Data)
	log.Println(a.Mode(false))
}

func TestArray(b *testing.T) {
	a := Arrayf64{}
	a.Init(Optionf64{
		Degree: 4,
	})
	populate(&a, 1000)
	log.Printf("%+v\n", a.Summary())
	a.Center(true)
	log.Printf("%+v\n", a.Summary())
	a.Reduce(true)
	log.Printf("%+v\n", a.Summary())
}

func BenchmarkArraySummary(b *testing.B) {
	for i := 1; i < 7; i++ {
		a := Arrayf64{}
		a.Init(Optionf64{
			Degree: 4,
		})

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
		for i := 1; i <= 6; i++ {
			a := Arrayf64{}
			a.Init(Optionf64{
				Degree: 2,
			})

			populate(&a, int(math.Pow(10, float64(i))))
			b.ResetTimer()
			b.Run(fmt.Sprintf("BenchmarkArrayApply_%s_10^%d", name, i), func(b *testing.B) {
				a.apply(funct, true)
			})
		}
	}
}
