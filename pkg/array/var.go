package array

import "math"

// Var computes the variance of the data array
// The structure stores both E[X^2] and E[X]^2 so it comes down to:
//		(array.Sum[1] -  math.Pow(array.Sum[0], 2)/array.Length) / (array.Length - 1)
//		(E[X^2] - E[X]^2)/(n - 1)
//
// Complexity:
// 	=	1 memory access + 1 float64-float64 substration + 1 math.Pow(float64, 2)
//		+ 1 memory access + 1 float64-int division + 1 float64-float64 division + 1 float64-int subtraction
//	=   2 memory access + 2 subtraction + 2 division + 1 math.Pow(float64,2)
//  =   2 memory access + 2 subtraction + 2 division + 1 63 bit shift
//  ~ 	O(1)
//
func (a *Arrayf64) Var() float64 {
	if a.Length > 0 {
		return (a.Sum[1] - math.Pow(a.Sum[0], 2)/a.Length) / (a.Length - 1)
	}
	return 0
}

// Stddev computes the standard deviation of the data array
// Algorithm:
//		√(VAR[X])
//
// Complexity:
// =	1 Var operation + 1 math.Sqrt(float64)
// ~	O(1) + O(1)
// ~	O(1)
//
func (a *Arrayf64) Stddev() float64 {
	return math.Sqrt(a.Var())
}

// IQR returns the Interquartile range of the data array
// Algorithm:
//		q3 - q1
//
// Complexity:
// =	2 quantile + 1 float64-float64 subtraction
// ~	2 * O(1) + O(1)
// ~	O(1)
//
func (a *Arrayf64) IQR() float64 {
	return a.Quantile(.75) - a.Quantile(.25)
}

// Midhinge returns the Midhinge of the data array
// Algorithm:
//		(q1 + q3)/2
//
// Complexity:
// =	2 quantile + 1 float64-float64 division + 1 float64-float64 addition
// ~	2 * O(1) + O(1) + O(1)
// ~	O(1)
//
func (a *Arrayf64) Midhinge() float64 {
	return (a.Quantile(.25) + a.Quantile(.75)) / 2
}

// Trimean returns the Trimean of the data array
// Algorithm:
//		(q2 + midhinge)/2
//
// Complexity:
// =	1 quantile + 1 midhinge +
//      1 float64-float64 division + 1 float64-float64 addition
// ~	O(1) + O(1) + O(1) + O(1)
// ~	O(1)
//
func (a *Arrayf64) Trimean() float64 {
	return (a.Quantile(.5) + a.Midhinge()) / 2
}
