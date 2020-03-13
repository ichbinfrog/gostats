package array

import "math"

// Quantile computes the quantiles of the given data array
// Since the array is sorted, the n-th element of the array is
// similarly the n-th smallest element. Accessing a given q quantile
// loops back to being a memory access at q * array.Length:
// Algorithm:
//		array.Data[floor(q * array.Length)]
//
// Complexity:
// = 	1 memory access + 1 float64-int cast + 1 math.Floor(float64) + 1 float64-float64 mult
// ~	O(1)
//
func (a *Arrayf64) Quantile(q float64) float64 {
	return a.Data[int(math.Floor(q*a.Length))]
}

// Median returns quantile(.5)c
func (a *Arrayf64) Median() float64 {
	return a.Quantile(.5)
}

// Min returns the first element of the sorted array
// Algorithm:
//		array.Data[0]
//
// Complexity:
// =	1 memory access
// ~	O(1)
//
func (a *Arrayf64) Min() float64 {
	if a.Length > 0 {
		return a.Data[0]
	}
	return 0
}

// Max returns the laster element of the sorted array
// Algorithm:
//		array.Data[array.Length - 1]
//
// Complexity:
// =	1 memory access + 1 float64-int subtraction
// ~	O(1)
//
func (a *Arrayf64) Max() float64 {
	if a.Length > 0 {
		return a.Data[int(a.Length)-1]
	}
	return 0
}
