package array

import (
	"math"
	"sort"
)

var (
	// AggregateMap stores the funciton associated with the unmutable
	// statistic aggregate. Unmutable here meaning that the applying
	// changes to the underlying array would require recomputing
	// the entire aggregate instead of the single change.
	AggregateMap map[string]Aggregate
)

// Aggregate groups iterative and summation function
// for a statistic aggreate.
type Aggregate struct {
	Iterative func(float64, float64) float64
	Summation func(bool) float64
}

func init() {
	AggregateMap = make(map[string]Aggregate)
}

// Optionf64 represents the Option group to select which
// variable to accelerate
type Optionf64 struct {
	Degree    int  `json:"degree"`
	Harmonic  bool `json:"harmonic"`
	Geometric bool `json:"geometric"`
}

// Arrayf64 is a statistics wrapper around an array of float64
type Arrayf64 struct {
	Option    Optionf64          `json:"options"`
	Length    float64            `json:"length"`
	Sum       []float64          `json:"sum"`
	Data      []float64          `json:"data"`
	Aggregate map[string]float64 `json:"aggregate"`
}

// Init allocates the Sum array with a given degree
func (a *Arrayf64) Init(opt Optionf64) {
	a.Option = opt
	a.Sum = make([]float64, opt.Degree)
	a.Aggregate = make(map[string]float64)

	if opt.Geometric {
		AggregateMap["geometric"] = Aggregate{
			Iterative: geometricAdd,
			Summation: a.GeometricMean,
		}
		a.Aggregate["geometric"] = 0.0
	}
	if opt.Harmonic {
		AggregateMap["harmonic"] = Aggregate{
			Iterative: harmonicAdd,
			Summation: a.HarmonicMean,
		}
		a.Aggregate["harmonic"] = 0.0
	}
}

// Insert inserts the value in the sorted array
// Algorithm:
// 	Update aggregates and length
// 	Update aggregate data
//	Find index where value should be inserted
//  Shift slice to [index + 1]
//  Insert array at [index]
//
// Complexity:
//		O(Aggregate update) + O(index find) + O(shift slice) + O(insert)
//		= O(1) + O(n) + O(n) + O(1)
//		= O(n)
//
func (a *Arrayf64) Insert(val float64) {
	for i := 0; i < a.Option.Degree; i++ {
		if i == 0 {
			a.Sum[i] += val
		} else {
			a.Sum[i] += math.Pow(val, float64(i+1))
		}
	}
	for k, f := range a.Aggregate {
		a.Aggregate[k] = AggregateMap[k].Iterative(f, val)
	}
	a.Length++

	index := sort.SearchFloat64s(a.Data, val)
	a.Data = append(a.Data, 0)
	copy(a.Data[index+1:], a.Data[index:])
	a.Data[index] = val
}

// InsertSlice inserts a slice of float64 value in the sorted array
// Algorithm:
//  Insert value at [index] for value in slice
//
// Complexity:
//		len(sl) * insert(val)
//		=> O(len(sl) * insert(val))
//		=> O(|sl|)
//
func (a *Arrayf64) InsertSlice(values []float64) {
	for _, val := range values {
		a.Insert(val)
	}
}

// At returns a pointer to the value at a given index
func (a *Arrayf64) At(index int) *float64 {
	if index <= int(a.Length) {
		return &a.Data[index]
	}
	return nil
}

func (a *Arrayf64) updateAggregates(old *float64, new float64) {
	for i := 0; i < a.Option.Degree; i++ {
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

func (a *Arrayf64) apply(f func(float64) float64, update bool) {
	for i := 0; i < int(a.Length); i++ {
		a.Change(i, f(a.Data[i]), update)
	}

	for k := range a.Aggregate {
		a.Aggregate[k] = AggregateMap[k].Summation(true)
	}
}

// DeepCopy returns a pointer of an exact copy of an array
func (a *Arrayf64) DeepCopy() *Arrayf64 {
	na := &Arrayf64{
		Option: a.Option,
		Length: a.Length,
	}
	na.Data = make([]float64, int(na.Length))
	na.Sum = make([]float64, a.Option.Degree)
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
