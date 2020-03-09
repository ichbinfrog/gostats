package array

import "math"

// Sigmoid returns the data set within -1 1
func (a *Arrayf64) Sigmoid(inplace bool) *Arrayf64 {
	if a.Length == 0 {
		return nil
	}
	if inplace {
		a.apply(func(v float64) float64 {
			return 1 / (1 + math.Exp(-v))
		}, true)
		return a
	}
	na := a.DeepCopy()
	na.apply(func(v float64) float64 {
		return 1 / (1 + math.Exp(-v))
	}, true)
	return na
}
