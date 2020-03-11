package dist

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"

	"github.com/ichbinfrog/statistics/pkg/util"
	"gonum.org/v1/gonum/mathext"
)

// Binomial represents the Binomial distribution
// Discreet probability distribution function as follows:
// 		X ~ B(n, p), n >= 0, 0 <= p <= 1
//		f(k,p) = {
//			C(n, k)p^k * q^(n - k)
//		}, k in [0, ..., n]
//
type Binomial struct {
	N, P, Q float64
}

// Generate creates one sample of the Bernouilli distribution
func (b *Binomial) Generate() float64 {
	// PRESS, William H., TEUKOLSKY, Saul A., VETTERLING, William T., et al. Numerical recipes in C. 1988.
	p := b.P
	if b.P > .5 {
		p = b.Q
	}

	// Direct method
	if b.N < 25 {
		bnl := 0.0
		for i := 1; i <= int(b.N); i++ {
			if rand.Float64() < p {
				bnl++
			}
		}
		if p != b.P {
			return b.N - bnl
		}
		return bnl
	}

	am := b.N * p
	// Direct Poisson method
	if am < 1.0 {
		g := math.Exp(-am)
		t := 1.0
		i := 0.0
		for {
			t *= rand.Float64()
			if t < g || i >= b.N {
				if p != b.P {
					return b.N - i
				}
				return i
			}
			i++
		}
	}

	// Rejection method
	sq := math.Sqrt(2.0 * am * p)
	oldg, _ := math.Lgamma(b.N + 1.0)
	plog := math.Log(p)
	pclog := math.Log(1.0 - p)

	for {
		var y, em float64
		for {
			y = math.Tan(math.Pi * rand.Float64())
			em = sq*y + am
			if em >= 0.0 && em < b.N+1 {
				break
			}
		}
		em = math.Floor(em)
		lg1, _ := math.Lgamma(em + 1.0)
		lg2, _ := math.Lgamma(b.N - em + 1.0)
		t := 1.2 * math.Sqrt(1.0+math.Pow(y, 2)) * math.Exp(oldg-lg1-lg2+em*plog+(b.N-em)*pclog)

		if rand.Float64() <= t {
			if p != b.P {
				return b.N - em
			}
			return em
		}
	}
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
func (b *Binomial) Init(n, p float64) error {
	if p < 0 || p > 1 || n <= 0 {
		return util.ErrBinomialParam
	}
	b.N, b.P, b.Q = n, p, 1-p
	return nil
}

// Domain returns the definition domain of the distribution
func (b *Binomial) Domain() (float64, float64) {
	return 0, float64(b.N)
}

// PMF returns the probability mass function value of a given k
func (b *Binomial) PMF(k float64) float64 {
	return float64(BinomialCoeff(int(b.N), int(k))) * math.Pow(b.P, k) * math.Pow(b.Q, float64(b.N)-k)
}

// CDF returns the Cumulative distribution function value of a given k
func (b *Binomial) CDF(k float64) float64 {
	if k < 0 {
		return 0
	}
	if k >= b.N {
		return 1
	}
	k = math.Floor(k)
	return mathext.RegIncBeta(b.N-k, k+1, b.Q)
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

// Summary returns a string summarising basic info about the distribution
func (b *Binomial) Summary() string {
	dbeg, dend := b.Domain()
	return fmt.Sprintf(`
	X ~ B(%f, %f)
		Domain:			{ %f , %f }
		Mean: 			%f
		Median (upper): %f
		Median (lower): %f
		Var: 			%f
		Skewness: 		%f
		Kurtosis:		%f
		FisherInfo:		%f
`, b.N, b.P, dbeg, dend, b.Mean(), b.Median(true), b.Median(false), b.Var(), b.Skewness(), b.Kurtosis(), b.FisherI())
}
