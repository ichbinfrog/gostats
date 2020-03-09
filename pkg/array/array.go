package array

import (
	"math"
	"sort"
)

// Arrayf64 is a statistics wrapper around an array of float64
type Arrayf64 struct {
	Degree int       `json:"degree"`
	Length float64   `json:"length"`
	Sum    []float64 `json:"sum"`
	Data   []float64 `json:"data"`
}

// Init allocates the Sum array with a given degree
func (a *Arrayf64) Init(m int) {
	if m > 1 {
		a.Sum = make([]float64, m)
		a.Degree = m
	}
}

// Insert inserts the value in the sorted array
// Algorithm:
// 	Updates aggregates and length
//	Finds index where value should be inserted
//  Shift slice to [index + 1]
//  Insert array at [index]
//
// Complexity:
//		O(Aggregate update) + O(index find) + O(shift slice) + O(insert)
//		= O(1) + O(n) + O(n) + O(1)
//		= O(n)
//
func (a *Arrayf64) Insert(val float64) {
	for i := 0; i < a.Degree; i++ {
		if i == 0 {
			a.Sum[i] += val
		} else {
			a.Sum[i] += math.Pow(val, float64(i+1))
		}
	}
	a.Length++

	index := sort.SearchFloat64s(a.Data, val)
	a.Data = append(a.Data, 0)
	copy(a.Data[index+1:], a.Data[index:])
	a.Data[index] = val
}

// At returns a pointer to the value at a given index
func (a *Arrayf64) At(index int) *float64 {
	if index <= int(a.Length) {
		return &a.Data[index]
	}
	return nil
}

func (a *Arrayf64) updateAggregates(old *float64, new float64) {
	for i := 0; i < a.Degree; i++ {
		a.Sum[i] = a.Sum[i] + math.Pow(new, float64(i+1)) - math.Pow(*old, float64(i+1))
	}
}

// Change modifies the value at a given index with a given value
func (a *Arrayf64) Change(index int, val float64, update bool) {
	if old := a.At(index); old != nil {
		if update {
			a.updateAggregates(old, val)
		}
		*old = val
	}
}

// Remove pops the data at the given index
func (a *Arrayf64) Remove(index int) {
	if old := a.At(index); old != nil {
		a.updateAggregates(old, 0)
		a.Data = append(a.Data[:index], a.Data[:index+1]...)
	}
}

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
// =	1 memory access + 1 float64-int substraction
// ~	O(1)
//
func (a *Arrayf64) Max() float64 {
	if a.Length > 0 {
		return a.Data[int(a.Length)-1]
	}
	return 0
}

// skewnessMeasure
type skewnessMeasure int8

const (
	// Yule is a quantile-based measure of skewness
	Yule skewnessMeasure = iota
	// PearsonSecond is a median-based measure of skewness
	PearsonSecond
)

// Skewness returns a skewness measure of the given data set
// Algorithm:
// 	Case MeasureType:
//		Yule: (q3 + q1 - 2M)/ (q3 - q1)
//		Pearson Second: 3 (E[X] - M)/σ
//
// Complexity:
//		max(O(Yule), O(Pearson Second))
//		= max(O(1), O(1))
//		= O(1)
//
func (a *Arrayf64) Skewness(s skewnessMeasure) float64 {
	switch s {
	case Yule:
		q3 := a.Quantile(.3)
		q1 := a.Quantile(.1)
		return (q3 + q1 - 2*a.Quantile(.2)) / (q3 - q1)
	case PearsonSecond:
		return 3 * (a.Mean() - a.Median()/a.Stddev())
	default:
		return 0
	}
}

// Summaryf64 is the pandas describe's equivalent structure
type Summaryf64 struct {
	Length float64 `json:"length"`
	Mean   float64 `json:"mean"`
	Stddev float64 `json:"stddev"`
	Min    float64 `json:"min"`
	Max    float64 `json:"max"`
	Median float64 `json:"median"`
	Q1     float64 `json:"q1"`
	Q3     float64 `json:"q3"`
}

// Summary returns the summary of the data set
func (a *Arrayf64) Summary() *Summaryf64 {
	return &Summaryf64{
		Length: a.Length,
		Mean:   a.Mean(),
		Stddev: a.Stddev(),
		Min:    a.Min(),
		Max:    a.Max(),
		Median: a.Median(),
		Q1:     a.Quantile(.25),
		Q3:     a.Quantile(.75),
	}
}

func (a *Arrayf64) apply(f func(float64) float64, update bool) {
	for i := 0; i < int(a.Length); i++ {
		a.Change(i, f(a.Data[i]), update)
	}
}

// DeepCopy returns a pointer of an exact copy of an array
func (a *Arrayf64) DeepCopy() *Arrayf64 {
	na := &Arrayf64{
		Degree: a.Degree,
		Length: a.Length,
	}
	copy(na.Data, a.Data)
	copy(na.Sum, a.Sum)
	return na
}

// Center centers the dataset around the mean
// If uncentered data is inserted the center operation
// the dataset would be effectively corrupted
func (a *Arrayf64) Center(inplace bool) *Arrayf64 {
	x := a.Mean()
	if inplace {
		a.apply(func(v float64) float64 {
			return v - x
		}, true)
		return nil
	}
	na := a.DeepCopy()
	na.apply(func(v float64) float64 {
		return v - x
	}, true)
	return na
}

// Reduce normalises the dataset around the stddev
// If non normalised data is inserted the center operation
// the dataset would be effectively corrupted
func (a *Arrayf64) Reduce(inplace bool) *Arrayf64 {
	s := a.Stddev()
	if inplace {
		a.apply(func(v float64) float64 {
			return v / s
		}, true)
		return nil
	}
	na := a.DeepCopy()
	na.apply(func(v float64) float64 {
		return v / s
	}, true)
	return na
}
