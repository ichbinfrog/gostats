package array

import "math"

// Mode gets the most common occurence in the data array
// Complexity: O(n)
//
func (a *Arrayf64) Mode() float64 {
	maxCount, maxVal := 0, math.Inf(-1)
	currCount, currVal := 0, math.Inf(-1)

	for _, v := range a.Data {
		if math.IsInf(currVal, -1) {
			currVal = v
			currCount = 1
		}

		if v == currVal {
			currCount++
		} else {
			if currCount > maxCount {
				maxCount = currCount
				maxVal = currVal
			}

			currVal = v
			currCount = 1
		}
	}

	if currCount > maxCount {
		maxCount = currCount
		maxVal = currVal
	}

	return maxVal
}
