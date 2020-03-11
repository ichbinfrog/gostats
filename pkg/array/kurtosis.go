package array

import "math"

// Kurtosis returns an unbiased estimation of the
// the given data set. Complexity: O(n)
func (a *Arrayf64) Kurtosis() float64 {
	mean := a.Mean()
	k4 := 0.0
	for i := 0; i < int(a.Length); i++ {
		k4 += math.Pow(a.Data[i]-mean, 4)
	}
	return k4/(a.Length*math.Pow(a.Var(), 2)) - 3
}
