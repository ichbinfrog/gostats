package array

import "math"

// Entropy returns the entropy of a given dataset
func (a *Arrayf64) Entropy() float64 {
	if centered := a.Center(false); centered != nil {
		if reduced := centered.Reduce(false); reduced != nil {
			entropy := 0.0
			for _, v := range reduced.Data {
				entropy += (v * math.Log(v))
			}
			return entropy
		}
	}
	return math.NaN()
}
