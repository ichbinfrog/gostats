package dist

import (
	"math"
	"math/big"
)

// Poisson represents the Poisson distribution
// Continuous probability distribution function as follows:
// 		X ~ P(λ),  λ in ] 0, +inf [
//		P(X = k) = (λ^k / k!)e^(-λ)
//
type Poisson struct {
	Lambda float64
}

// Init intialises a Bernouilli distribution
func (p *Poisson) Init(lambda float64) {
	if lambda < 0 {
		panic("")
	}
	p.Lambda = lambda
}

// Factorial computes the binomial coeff with the given n, k
func Factorial(n int) int64 {
	if n == 0 || n == 1 {
		return 1
	}
	r1 := big.NewInt(1)
	for i := 2; i <= n; i++ {
		r1.Mul(r1, big.NewInt(int64(i)))
	}
	return r1.Int64()
}

// Domain returns the definition domain of the distribution
func (p *Poisson) Domain() (float64, float64) {
	return 0, math.Inf(0)
}

// PMF returns the probability mass function value of a given k
func (p *Poisson) PMF(k float64) float64 {
	return math.Pow(p.Lambda, k) * math.Exp(-p.Lambda) / float64(Factorial(int(k)))
}

// CDF returns the Cumulative distribution function value of a given k
func (p *Poisson) CDF(k float64) float64 {
	sum := 0.0
	for i := 0; i < int(math.Floor(p.Lambda)); i++ {
		sum += math.Pow(p.Lambda, float64(i)) / float64(Factorial(i))
	}
	sum *= math.Exp(-p.Lambda)
	return sum
}

// Mean returns the mean of the distribution
func (p *Poisson) Mean() float64 {
	return p.Lambda
}

// Median returns the median of the distribution
func (p *Poisson) Median(upper bool) float64 {
	return math.Ceil(p.Lambda + 1/3 - 0.02*p.Lambda)
}

// Var returns the variance of the distribution
func (p *Poisson) Var() float64 {
	return p.Lambda
}

// Skewness returns the Pearson's moment coefficient of skewness of the distribution
func (p *Poisson) Skewness() float64 {
	return math.Pow(p.Lambda, -1/2)
}

// Kurtosis returns the Kurtosis of the distribution
func (p *Poisson) Kurtosis() float64 {
	return 1 / p.Lambda
}

// Moment returns the t-th moment of the distribution
func (p *Poisson) Moment(t float64) float64 {
	return math.Exp(p.Lambda * (math.Exp(t) - 1))
}

// FisherI returns the Fisher Information of the distribution
func (p *Poisson) FisherI() float64 {
	return 1 / p.Lambda
}