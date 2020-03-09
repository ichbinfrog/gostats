package dist

import (
	"math"
	"math/rand"
)

// Bernouilli represents the Bernouilli distribution
// Discreet probability distribution function as follows:
// 		X ~ B(p), 0 <= p <= 1
//		f(k,p) = {
//			if k = 1, p
//			if k = 0, 1 - p
//		}, k in [0 , 1]
//
type Bernouilli struct {
	P float64
	Q float64
}

// Init intialises a Bernouilli distribution
func (b *Bernouilli) Init(p float64) {
	if p < 0 || p > 1 {
		panic("")
	}
	b.P, b.Q = p, 1-p
}

// Domain returns the definition domain of the distribution
func (b *Bernouilli) Domain() (float64, float64) {
	return 0, 1
}

// Generate creates one sample of the Bernouilli distribution
func (b *Bernouilli) Generate() bool {
	return rand.Float64() < b.P
}

// PMF returns the probability mass function value of a given k
func (b *Bernouilli) PMF(k float64) float64 {
	if k == 0 {
		return b.Q
	}
	return b.P
}

// CDF returns the Cumulative distribution function value of a given k
func (b *Bernouilli) CDF(k float64) float64 {
	if k < 0 {
		return 0
	}
	if k >= 1 {
		return 1
	}
	return b.Q
}

// Mean returns the mean of the distribution
func (b *Bernouilli) Mean() float64 {
	return b.P
}

// Median returns the median of the distribution
func (b *Bernouilli) Median() float64 {
	if b.P < .5 {
		return 0
	}
	if b.P > .5 {
		return 1
	}
	return .5
}

// Var returns the variance of the distribution
func (b *Bernouilli) Var() float64 {
	return b.P * b.Q
}

// Skewness returns the Pearson's moment coefficient of skewness of the distribution
func (b *Bernouilli) Skewness() float64 {
	return (b.Q - b.P) / math.Sqrt(b.P*(b.Q))
}

// Kurtosis returns the Kurtosis of the distribution
func (b *Bernouilli) Kurtosis() float64 {
	return (1 - 6*b.P*b.Q) / (b.P * b.Q)
}

// Entropy returns the Entropy of the distribution
func (b *Bernouilli) Entropy() float64 {
	return -b.Q*math.Log(b.Q) - b.P*math.Log(b.P)
}

// Moment returns the t-th moment of the distribution
func (b *Bernouilli) Moment(t float64) float64 {
	return b.Q + b.P*math.Exp(t)
}

// FisherI returns the Fisher Information of the distribution
func (b *Bernouilli) FisherI() float64 {
	return 1 / (b.P * b.Q)
}
