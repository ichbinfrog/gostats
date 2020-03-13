package dist

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mathext"
)

// Chisq represents the Chi squared distribution
// Continuous probability distribution function as follows:
// 		X ~ χ(k), k >= 0
//
type Chisq struct {
	Degree float64
}

// Generate creates one sample of the Bernoulli distribution
func (c *Chisq) Generate() float64 {
	g := Gamma{}
	if err := g.Init(c.Degree/2, .5); err != nil {
		return math.NaN()
	}
	return g.Generate()
}

// PMF returns the probability mass function value of a given k
func (c *Chisq) PMF(x float64) float64 {
	lg, _ := math.Lgamma(c.Degree / 2)
	return (math.Pow(x, c.Degree/2-1) * math.Exp(-x/2)) / (math.Pow(2, c.Degree/2) * lg)
}

// CDF returns the Cumulative distribution function value of a given k
func (c *Chisq) CDF(x float64) float64 {
	lg, _ := math.Lgamma(c.Degree / 2)
	return mathext.GammaIncReg(c.Degree/2, x/2) / lg
}

// Mean returns the mean of the distribution
func (c *Chisq) Mean() float64 {
	return c.Degree
}

// Median returns the median of the distribution
func (c *Chisq) Median() float64 {
	return c.Degree * math.Pow(1-2/(9*c.Degree), 3)
}

// Var returns the variance of the distribution
func (c *Chisq) Var() float64 {
	return 2 * c.Degree
}

// Skewness returns the Pearson's moment coefficient of skewness of the distribution
func (c *Chisq) Skewness() float64 {
	return math.Sqrt(c.Degree - 2.0)
}

// Kurtosis returns the Kurtosis of the distribution
func (c *Chisq) Kurtosis() float64 {
	return 12 / c.Degree
}

// Moment returns the t-th moment of the distribution
func (c *Chisq) Moment(t float64) float64 {
	if t < 1/2 {
		return math.Pow(1-2*t, -c.Degree/2)
	}
	return math.NaN()
}

// Summary returns a string summarising basic info about the distribution
func (c *Chisq) Summary() string {
	return fmt.Sprintf(`
	X ~ χ(%f)
		Mean: 		%f
		Median: 	%f
		Var: 		%f
		Skewness: 	%f
		Kurtosis:	%f
`, c.Degree, c.Mean(), c.Median(), c.Var(), c.Skewness(), c.Kurtosis())
}
