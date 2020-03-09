package dist

import (
	"math"
)

// Polya represents the Polya distribution
// Discreet probability distribution function as follows:
// 		X ~ NB(r, p), r > 0, 0 <= p <= 1
//		f(k,p) = {
//			C(k + r - 1, k)q^r * p^k
//		}, k in [0, ..., n]
//
type Polya struct {
	R, P, Q float64
}

// Init intialises a Bernouilli distribution
func (p *Polya) Init(r, prob float64) {
	if prob < 0 || prob > 1 || r < 0 {
		panic("")
	}
	p.R, p.P, p.Q = r, prob, 1-prob
}

// Domain returns the definition domain of the distribution
func (p *Polya) Domain() (float64, float64) {
	return 0, math.Inf(0)
}

// PMF returns the probability mass function value of a given k
func (p *Polya) PMF(k int) float64 {
	return float64(BinomialCoeff(k+int(p.R)-1, k)) * math.Pow(p.Q, p.R) * math.Pow(p.P, float64(k))
}

// Mean returns the mean of the distribution
func (p *Polya) Mean() float64 {
	return (p.P * float64(p.R)) / p.Q
}

// Median returns the median of the distribution
func (p *Polya) Median(upper bool) float64 {
	if p.R > 1 {
		return math.Floor((p.P * (p.R - 1)) / p.Q)
	}
	return 0
}

// Var returns the variance of the distribution
func (p *Polya) Var() float64 {
	return p.P * p.R / math.Pow(p.Q, 2)
}

// Skewness returns the Pearson's moment coefficient of skewness of the distribution
func (p *Polya) Skewness() float64 {
	return (1 - p.P) / math.Sqrt(p.P*p.R)
}

// Kurtosis returns the Kurtosis of the distribution
func (p *Polya) Kurtosis() float64 {
	return (6 / p.R) + math.Pow(p.Q, 2)/(p.P*p.R)
}

// Moment returns the t-th moment of the distribution
func (p *Polya) Moment(t float64) float64 {
	if t < -math.Log(p.P) {
		return math.Pow(p.Q/(1-p.P*math.Exp(t)), p.R)
	}
	return math.NaN()
}

// FisherI returns the Fisher Information of the distribution
func (p *Polya) FisherI() float64 {
	return p.R / (p.P * math.Pow(p.Q, 2))
}
