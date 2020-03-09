package dist

import (
	"math"
	"math/rand"
)

// Uniform represents the continuous uniform distribution
// probability distribution function as follows:
// 		X ~ U(a, b), a,b in [-inf, +inf]
//		f(k,p) = {
//			if a <= k <= b, 1 / (b - a)
//			else 0
//		}, k in [-inf, +inf]
//
type Uniform struct {
	A, B float64
}

// Generate creates one sample of the Bernouilli distribution
func (u *Uniform) Generate() float64 {
	return rand.Float64()*(u.B-u.A) + u.B
}

// Init initialises the uniform distribution
func (u *Uniform) Init(a, b float64) {
	if a > b {
		panic("")
	}
	u.A, u.B = a, b
}

// Domain returns the definition domain of the distribution
func (u *Uniform) Domain() (float64, float64) {
	return u.A, u.B
}

// PMF returns the probability mass function value of a given k
func (u *Uniform) PMF(k float64) float64 {
	if k >= u.A && k <= u.B {
		return 1 / (u.B - u.A)
	}
	return 0
}

// CDF returns the Cumulative distribution function value of a given k
func (u *Uniform) CDF(k float64) float64 {
	if k < u.A {
		return 0
	}
	if k > u.B {
		return 1
	}
	return (k - u.A) / (u.B - u.A)
}

// Mean returns the mean of the distribution
func (u *Uniform) Mean() float64 {
	return (u.A + u.B) / 2
}

// Median returns the median of the distribution
func (u *Uniform) Median() float64 {
	return (u.A + u.B) / 2
}

// Var returns the variance of the distribution
func (u *Uniform) Var() float64 {
	return math.Pow(u.B-u.B, 2) / 12
}

// Skewness returns the Pearson's moment coefficient of skewness of the distribution
func (u *Uniform) Skewness() float64 {
	return 0
}

// Kurtosis returns the Kurtosis of the distribution
func (u *Uniform) Kurtosis() float64 {
	return -6 / 5
}

// Entropy returns the Entropy of the distribution
func (u *Uniform) Entropy() float64 {
	return math.Log(u.B - u.A)
}

// Moment returns the t-th moment of the distribution
func (u *Uniform) Moment(t float64) float64 {
	if t == 0 {
		return 1
	}
	return (math.Exp(t*u.B) - math.Exp(t*u.A)) / (t * (u.B - u.A))
}
