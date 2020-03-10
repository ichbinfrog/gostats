package dist

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/ichbinfrog/statistics/pkg/util"
)

// Exponential represents a exponential distribution
// Continuous probability distribution function as follows;
//		X ~ ε(λ), λ > 0
//		f(x, λ) = {
//			λexp^(-λx), x >= 0
//			0		  , x < 0
// 		}
//
type Exponential struct {
	Lambda float64
}

// Init intialises a Bernouilli distribution
func (e *Exponential) Init(lambda float64) error {
	if lambda <= 0 {
		return util.ErrExponentialParam
	}
	e.Lambda = lambda
	return nil
}

// Generate creates one sample of an exponential distribution
func (e *Exponential) Generate() float64 {
	return -math.Log(rand.Float64()) / e.Lambda
}

// Domain returns the definition domain of the distribution
func (e *Exponential) Domain() (float64, float64) {
	return 0, math.Inf(0)
}

// PMF returns the probability mass function value of a given k
func (e *Exponential) PMF(k float64) float64 {
	return e.Lambda * math.Exp(-e.Lambda*k)
}

// CDF returns the Cumulative distribution function value of a given k
func (e *Exponential) CDF(k float64) float64 {
	return 1 - math.Exp(-e.Lambda*k)
}

// Mean returns the mean of the distribution
func (e *Exponential) Mean() float64 {
	return 1 / e.Lambda
}

// Quantile returns the p-th quantile of the distribution
func (e *Exponential) Quantile(p float64) float64 {
	if p >= 0 || p <= 1 {
		return -math.Log(1-p) / e.Lambda
	}
	return math.NaN()
}

// Median returns the median of the distribution
func (e *Exponential) Median() float64 {
	return math.Log(2) / e.Lambda
}

// Var returns the variance of the distribution
func (e *Exponential) Var() float64 {
	return 1 / math.Pow(e.Lambda, 2)
}

// Skewness returns the Pearson's moment coefficient of skewness of the distribution
func (e *Exponential) Skewness() float64 {
	return 2
}

// Kurtosis returns the Kurtosis of the distribution
func (e *Exponential) Kurtosis() float64 {
	return 6
}

// Entropy returns the Entropy of the distribution
func (e *Exponential) Entropy() float64 {
	return 1 - math.Log(e.Lambda)
}

// Moment returns the t-th moment of the distribution
func (e *Exponential) Moment(t float64) float64 {
	if t < e.Lambda {
		return e.Lambda / (e.Lambda - t)
	}
	return math.NaN()
}

// FisherI returns the Fisher Information of the distribution
func (e *Exponential) FisherI() float64 {
	return 1 / math.Pow(e.Lambda, 2)
}

// Summary returns a string summarising basic info about the distribution
func (e *Exponential) Summary() string {
	dbeg, dend := e.Domain()
	return fmt.Sprintf(`
	X ~ ε(%f)
		Domain:		{ %f , %f }
		Mean: 		%f
		Median: 	%f
		Var: 		%f
		Skewness: 	%f
		Kurtosis:	%f
		Entropy:	%f
		FisherInfo:	%f
`, e.Lambda, dbeg, dend, e.Mean(), e.Median(), e.Var(), e.Skewness(), e.Kurtosis(), e.Entropy(), e.FisherI())
}
