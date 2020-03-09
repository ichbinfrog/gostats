package array

import (
	"math"
)

// Mean computes the mean of the data array
// Since the structure stores the sum value of all the inserted data as well
// as the length, it essentially comes down to:
// 		array.Sum[0] / float64(array.Length)
//		Σ(i = 0; i < n, i++)(x_i) / n
//
// Complexity:
// 		1 memory access + 1 float64 division + 1 int->float64 conversion ~ O(1)
//
func (a *Arrayf64) Mean() float64 {
	if a.Length > 0 {
		return a.Sum[0] / float64(a.Length)
	}
	return 0
}

func geometricAdd(agg float64, val float64) float64 {
	if agg == 0 {
		return val
	}
	return agg * val
}

// GeometricMean returns the geometric mean of the data set
// Algorithm:
//			Π(i = 0; i < n; i++)(x_i) / n
// Complexity:
//		Recomputation: O(n)
//		Iterative computation: O(1)
//
func (a *Arrayf64) GeometricMean(recompute bool) float64 {
	if a.Option.Geometric {
		if a.Option.Geometric {
			if a.Length > 0 {
				fact := 0.0
				for _, v := range a.Data {
					fact = geometricAdd(fact, v)
					if fact == 0 {
						return 0
					}
				}
				return math.Pow(fact, 1/a.Length)
			}
		}
		return math.Pow(a.Aggregate["geometric"], 1/a.Length)
	}
	return math.NaN()
}

func harmonicAdd(agg float64, val float64) float64 {
	if val <= 0 {
		return math.NaN()
	}
	return agg + (1 / val)
}

// HarmonicMean returns the harmonic mean of the data set
// Algorithm:
//			Σ(i = 0; i < n; i++)(1 / x_i) / n
// Complexity:
//		Recomputation: O(n)
//		Iterative computation: O(1)
//
func (a *Arrayf64) HarmonicMean(recompute bool) float64 {
	if a.Option.Harmonic {
		if recompute {
			if a.Length > 0 {
				harm := 0.0
				for _, v := range a.Data {
					harm = harmonicAdd(harm, v)
					if math.IsNaN(harm) {
						return harm
					}
				}
				return harm / a.Length
			}
		}
		return a.Aggregate["harmonic"] / a.Length
	}
	return math.NaN()
}
