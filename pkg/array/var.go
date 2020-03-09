package array

import "math"

// Var computes the variance of the data array
// The structure stores both E[X^2] and E[X]^2 so it comes down to:
//		(array.Sum[1] -  math.Pow(array.Sum[0], 2)/array.Length) / (array.Length - 1)
//		(E[X^2] - E[X]^2)/(n - 1)
//
// Complexity:
// 	=	1 memory access + 1 float64-float64 substration + 1 math.Pow(float64, 2)
//		+ 1 memory access + 1 float64-int division + 1 float64-float64 division + 1 float64-int substraction
//	=   2 memory access + 2 substraction + 2 division + 1 math.Pow(float64,2)
//  =   2 memory access + 2 substraction + 2 division + 1 63 bit shift
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
