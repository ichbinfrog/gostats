package dist

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/ichbinfrog/statistics/pkg/util"
)

// Triangular represents the discreet triangular distribution
// probability distribution function as follows:
// 		X ~ T(a, b), a,b,c in [-inf, +inf], a <= c <= b
//		f(k,p) = {
//			if a <= k <= c, 2(k - a)/(b - a)(c - a)
//			if c <  k <  b, 2(b - k)/(b - a)(b - c)
//			else 0
//		}, k in [-inf, +inf]
//
type Triangular struct {
	A, B, C float64
}

// Init intialises a Bernoulli distribution
func (t *Triangular) Init(a, b, c float64) error {
	if b >= c && c >= a && b > a {
		t.A, t.B, t.C = a, b, c
		return nil
	}
	return util.ErrTriangularParam
}

// Generate creates one sample of the Triangular distribution
func (t *Triangular) Generate() float64 {
	u := rand.Float64()
	if 0 < u && u < t.CDF(u) {
		return t.A + math.Sqrt(u*(t.B-t.A)*(t.C-t.A))
	}
	return t.B - math.Sqrt((1-u)*(t.B-t.A)*(t.B-t.C))
}

// Domain returns the definition domain of the distribution
func (t *Triangular) Domain() (float64, float64) {
	return t.A, t.B
}

// PMF returns the probability mass function value of a given k
func (t *Triangular) PMF(k float64) float64 {
	if k < t.A {
		return 0
	}
	if t.A <= k && k <= t.C {
		return 2 * (k - t.A) / ((t.B - t.A) * (t.C - t.A))
	}
	if k == t.C {
		return 2 * (t.B - t.A)
	}
	if t.C < k && k <= t.B {
		return 2 * (t.B - k) / ((t.B - t.A) * (t.B - t.C))
	}
	return 1
}

// CDF returns the Cumulative distribution function value of a given k
func (t *Triangular) CDF(k float64) float64 {
	if k <= t.A {
		return 0
	}
	if t.A < k && k <= t.C {
		return math.Pow(k-t.A, 2) / ((t.B - t.A) * (t.B - t.C))
	}
	if t.C < k && k <= t.B {
		return 1 - math.Pow(t.B-k, 2)/((t.B-t.A)*(t.B-t.C))
	}
	return 1
}

// Mean returns the mean of the distribution
func (t *Triangular) Mean() float64 {
	return (t.A + t.B + t.C) / 3
}

// Median returns the median of the distribution
func (t *Triangular) Median() float64 {
	if t.C >= (t.A+t.B)/2 {
		return t.A + math.Sqrt(((t.B-t.A)/(t.C-t.A))/2)
	}
	return t.B - math.Sqrt(((t.B-t.A)/(t.C-t.A))/2)
}

// Var returns the variance of the distribution
func (t *Triangular) Var() float64 {
	return (math.Pow(t.A, 2) + math.Pow(t.B, 2) + math.Pow(t.C, 2) - t.A*t.B - t.A*t.C - t.B*t.C) / 18
}

// Skewness returns the Pearson's moment coefficient of skewness of the distribution
func (t *Triangular) Skewness() float64 {
	return (math.Sqrt(2) * (t.A + t.B - 2*t.C) * (2*t.A - t.B - t.C) * (t.A - 2*t.B + t.C)) / (5 * math.Pow((math.Pow(t.A, 2)+math.Pow(t.B, 2)+math.Pow(t.C, 2)-t.A*t.B-t.A*t.C-t.B*t.C), 3/2))
}

// Kurtosis returns the Kurtosis of the distribution
func (t *Triangular) Kurtosis() float64 {
	return -3 / 5
}

// Entropy returns the Entropy of the distribution
func (t *Triangular) Entropy() float64 {
	return 1/2 + math.Log((t.B-t.A)/2)
}

// Moment returns the t-th moment of the distribution
func (t *Triangular) Moment(k float64) float64 {
	return 2 * ((t.B-t.C)*math.Exp(t.A*k) - (t.B-t.A)*math.Exp(t.C*k) + (t.C-t.A)*math.Exp(t.B*k)) / ((t.B - t.A) * (t.C - t.A) * (t.B - t.C) * math.Pow(k, 2))
}

// Summary returns a string summarising basic info about the distribution
func (t *Triangular) Summary() string {
	dbeg, dend := t.Domain()
	return fmt.Sprintf(`
	X ~ B(%f, %f, %f)
		Domain:		{ %f , %f }
		Mean: 		%f
		Median: 	%f
		Var: 		%f
		Skewness: 	%f
		Kurtosis:	%f
		Entropy:	%f
`, t.A, t.B, t.C, dbeg, dend, t.Mean(), t.Median(), t.Var(), t.Skewness(), t.Kurtosis(), t.Entropy())
}
