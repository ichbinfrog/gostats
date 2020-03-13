package array

import "math"

// Mode gets the most common occurrence in the data array
// Complexity: O(n) (due to having to loop around the entire dataset)
// Possible optimisation :
// 		// change mode during insert
//		// check max_mode stored in the array struct
//
//		if max_mode.value == math.NaN() {
//			max_mode.value = value
//			max_mode.count = 1
// 		}
//		if max_mode.value == value inserted {
//			max_mode.count ++
//		} else {
//			current_mode = find index of the inserted
//			current_mode.count = count(amount of subsequent same value)
//			if current_mode.count > max_mode.count {
//				max_mode = current_mode
// 		}
//
//		The insert operation could jump up to O(n) in the case when the entire
//		datasest comprises of a single value. Assuming a large array should not contain
//		only the same value, the complexity would be around
//			max(O(mode count) + O(insert)) ~ max(O(card(mode_seq)), nlog(n))
//
//		O(1) can only be achieved by profoundly modifiying the underlying data structure
//		in order to maintain the count for each distinct value.	Although possible with
//			[ ( distinct_v1 , count_v1 ), ( distinct_v2, count_v2 ), ... ]
//
//		Any iteration through the array would no longer benefit from Go's iteration optimisation,
// 		and it would make Quantile and Insert operations ~ O(n)
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
