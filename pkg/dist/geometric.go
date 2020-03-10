package dist

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/ichbinfrog/statistics/pkg/util"
)

// Geometric represents a geometric distribution
// Discreet probability distribution function as follows:
//		X ~ G(p), 0 <= p <= 1
//		P(X = k) = q^(k - 1)p
//
type Geometric struct {
	P, Q float64
}

// Init intialises a geometric distribution
func (g *Geometric) Init(p float64) error {
	if p < 0 || p > 1 {
		return util.ErrGeometricParam
	}
	g.P, g.Q = p, 1-p
	return nil
}

// Generate creates one sample of a geometric distribution
func (g *Geometric) Generate() float64 {
	return math.Floor(math.Log(rand.Float64()) / math.Log(g.Q))
}

// Domain returns the definition domain of the distribution
func (g *Geometric) Domain() (float64, float64) {
	return 0, math.Inf(0)
}

// PMF returns the probability mass function value of a given k
func (g *Geometric) PMF(k float64) float64 {
	return math.Pow(g.Q, k-1) * g.P
}

// CDF returns the Cumulative distribution function value of a given k
func (g *Geometric) CDF(k float64) float64 {
	return 1 - math.Pow(g.Q, k)
}

// Mean returns the mean of the distribution
func (g *Geometric) Mean() float64 {
	return 1 / g.P
}

// Median returns the median of the distribution
func (g *Geometric) Median() float64 {
	return math.Floor(-1 / (math.Log(g.Q)))
}

// Var returns the variance of the distribution
func (g *Geometric) Var() float64 {
	return g.Q / math.Pow(g.P, 2)
}

// Skewness returns the Pearson's moment coefficient of skewness of the distribution
func (g *Geometric) Skewness() float64 {
	return (2 - g.P) / math.Sqrt(g.Q)
}

// Kurtosis returns the Kurtosis of the distribution
func (g *Geometric) Kurtosis() float64 {
	return 6 + math.Pow(g.P, 2)/g.Q
}

// Entropy returns the Entropy of the distribution
func (g *Geometric) Entropy() float64 {
	return (-(g.Q)*math.Log(g.Q) - g.P*math.Log(g.P)) / g.P
}

// Moment returns the t-th moment of the distribution
func (g *Geometric) Moment(t float64) float64 {
	return (g.P * math.Exp(t)) / (1 - g.Q*math.Exp(t))
}

// Summary returns a string summarising basic info about the distribution
func (g *Geometric) Summary() string {
	dbeg, dend := g.Domain()
	return fmt.Sprintf(`
	X ~ G(%f)
		Domain:			[ %f , %f [
		Mean: 			%f
		Median:			%f
		Var: 			%f
		Skewness: 		%f
		Entropy:		%f
		Kurtosis:		%f
`, g.P, dbeg, dend, g.Mean(), g.Median(), g.Var(), g.Skewness(), g.Kurtosis(), g.Entropy())
}
