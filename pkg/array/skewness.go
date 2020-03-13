package array

// skewnessMeasure
type skewnessMeasure int8

const (
	// Yule is a quantile-based measure of skewness
	Yule skewnessMeasure = iota
	// PearsonFirst is a mode-based measure of skewness
	PearsonFirst
	// PearsonSecond is a median-based measure of skewness
	PearsonSecond
)

// Skewness returns a skewness measure of the given data set
// Algorithm:
// 	Case MeasureType:
//		Yule: (q3 + q1 - 2M)/ (q3 - q1)
//		Pearson Second: 3 (E[X] - M)/σ
//		Pearson First: (E[X] - Mode)/σ
// Complexity:
//		max(O(Yule), O(Pearson Second), O(Pearson First))
//		= max(O(1), O(1), O(n))
//		= O(n)
//
func (a *Arrayf64) Skewness(s skewnessMeasure) float64 {
	switch s {
	case Yule:
		q3 := a.Quantile(.3)
		q1 := a.Quantile(.1)
		return (q3 + q1 - 2*a.Quantile(.2)) / (q3 - q1)
	case PearsonSecond:
		return 3 * (a.Mean() - a.Median()/a.Stddev())
	case PearsonFirst:
		return (a.Mean() - a.Mode(false)) / a.Stddev()
	default:
		return 0
	}
}
