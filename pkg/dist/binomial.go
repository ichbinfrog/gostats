package dist

import (
	"math"
	"math/big"
	"math/rand"
)

// Binomial represents the Binomial distribution
// Discreet probability distribution function as follows:
// 		X ~ B(n, p), n >= 0, 0 <= p <= 1
//		f(k,p) = {
//			C(n, k)p^k * q^(n - k)
//		}, k in [0, ..., n]
//
type Binomial struct {
	N    int
	P, Q float64
}

// Generate creates one sample of the Bernouilli distribution
func (b *Binomial) Generate() bool {
	return rand.Float64() < b.P
}

// BinomialCoeff computes the binomial coeff with the given n, k
func BinomialCoeff(n, k int) int64 {
	if k == 1 || n == k {
		return 1
	}
	r1, r2, r3 := big.NewInt(1), big.NewInt(1), big.NewInt(1)

	for i := 2; i <= n; i++ {
		r1.Mul(r1, big.NewInt(int64(i)))
		if i == n-k && r2.Cmp(big.NewInt(0)) != 0 {
			r2 = r2.Set(r1)
		}
		if i == k && r3.Cmp(big.NewInt(0)) != 0 {
			r3 = r3.Set(r1)
		}
	}

	return r1.Div(r1, r2.Mul(r2, r3)).Int64()
}

// Init intialises a Bernouilli distribution
func (b *Binomial) Init(n int, p float64) {
	if p < 0 || p > 1 || n < 0 {
		panic("")
	}
	b.N, b.P, b.Q = n, p, 1-p
}

// Domain returns the definition domain of the distribution
func (b *Binomial) Domain() (float64, float64) {
	return 0, float64(b.N)
}

// PMF returns the probability mass function value of a given k
func (b *Binomial) PMF(k float64) float64 {
	return float64(BinomialCoeff(b.N, int(k))) * math.Pow(b.P, k) * math.Pow(b.Q, float64(b.N)-k)
}

// Mean returns the mean of the distribution
func (b *Binomial) Mean() float64 {
	return b.P * float64(b.N)
}

// Median returns the median of the distribution
func (b *Binomial) Median(upper bool) float64 {
	if upper {
		return math.Ceil(b.Mean())
	}
	return math.Floor(b.Mean())
}

// Var returns the variance of the distribution
func (b *Binomial) Var() float64 {
	return b.P * b.Q * float64(b.N)
}

// Skewness returns the Pearson's moment coefficient of skewness of the distribution
func (b *Binomial) Skewness() float64 {
	return (b.Q - b.P) / math.Sqrt(b.Var())
}

// Kurtosis returns the Kurtosis of the distribution
func (b *Binomial) Kurtosis() float64 {
	return (1 - 6*b.P*b.Q) / (b.Var())
}

// Moment returns the t-th moment of the distribution
func (b *Binomial) Moment(t float64) float64 {
	return math.Pow(b.Q+b.P*math.Exp(t), float64(b.N))
}

// FisherI returns the Fisher Information of the distribution
func (b *Binomial) FisherI() float64 {
	return float64(b.N) / (b.P * b.Q)
}
