package array

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
