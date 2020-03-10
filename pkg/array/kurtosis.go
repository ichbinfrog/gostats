package array

import (
	"math"
)

// Kurtosis returns an unbiased estimation of the
// the given data set. Complexity: O(n)
func (a *Arrayf64) Kurtosis() float64 {
	n := a.Length
	mean := a.Mean()
	k2, k4 := 0.0, 0.0
	for i := 0; i < int(a.Length); i++ {
		k4 += math.Pow(a.Data[i]-mean, 4)
		k2 += math.Pow(a.Data[i]-mean, 2)
	}
	return ((n+1)*n)/((n-1)*(n-2)*(n-3))*(k4/math.Pow(k2, 2)) - 3*math.Pow(n-1, 2)/((n-2)*(n-3))
}
