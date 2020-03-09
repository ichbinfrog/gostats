package dist

import (
	"math"
)

// Geometric represents a geometric distribution
// Discreet probability distribution function as follows:
//		X ~ G(p), 0 <= p <= 1
//		P(X = k) = q^(k - 1)p
//
type Geometric struct {
	P, Q float64
}

// Init intialises a Bernouilli distribution
func (g *Geometric) Init(p float64) {
	if p < 0 || p > 1 {
		panic("")
	}
	g.P, g.Q = p, 1-p
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
