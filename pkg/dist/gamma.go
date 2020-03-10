package dist

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/ichbinfrog/statistics/pkg/util"
	"gonum.org/v1/gonum/mathext"
)

// Gamma represents a gamme distribution
// Continuous distribution function as follows:
//		X ~	Γ(α, β), α > 0, β > 0
//
//		f(x,α,β) = (β^α*x^(α-1)*exp(-β*x))/Γ(α)
//
type Gamma struct {
	Alpha, Beta float64
}

// Generate creates one sample of the Gamma distribution
func (g *Gamma) Generate() float64 {
	// Direct method
	if g.Beta < 6 {
		x := 1.0
		for i := 0.0; i < g.Beta; i++ {
			x *= rand.Float64()
		}
		return -math.Log(x)
	}

	// Rejection method
	var e float64
	for {
		var y, x, s, am float64
		for {
			var v1, v2 float64
			for {
				v1 = rand.Float64()
				v2 = 2.0*rand.Float64() - 1.0
				if math.Pow(v1, 2)+math.Pow(v2, 2) <= 1.0 {
					break
				}
			}
			y = v2 / v1
			am = g.Beta - 1
			s = math.Sqrt(2.0*am + 1.0)
			x = s*y + am
			if x > 0.0 {
				break
			}
		}
		e = (1.0 + math.Pow(y, 2)) / math.Exp(am*math.Log(x/am)-s*y)
		if rand.Float64() <= e {
			return x
		}
	}
}

// Init intialises a Bernouilli distribution
func (g *Gamma) Init(alpha, beta float64) error {
	if alpha <= 0 || beta <= 0 {
		return util.ErrGammaParam
	}
	g.Alpha, g.Beta = alpha, beta
	return nil
}

// Domain returns the definition domain of the distribution
func (g *Gamma) Domain() (float64, float64) {
	return 0, math.Inf(0)
}

// PMF returns the probability mass function value of a given k
func (g *Gamma) PMF(x float64) float64 {
	lg, _ := math.Lgamma(g.Alpha)
	return (math.Pow(g.Beta, g.Alpha) * math.Exp(-g.Beta*x) * math.Pow(x, g.Alpha-1)) / lg
}

// CDF returns the Cumulative distribution function value of a given k
func (g *Gamma) CDF(x float64) float64 {
	if x < 0 {
		return 0
	}
	return mathext.GammaIncReg(g.Alpha, x*g.Beta)
}

// Mean returns the mean of the distribution
func (g *Gamma) Mean() float64 {
	return g.Alpha / g.Beta
}

// Var returns the variance of the distribution
func (g *Gamma) Var() float64 {
	if g.Alpha >= 1 {
		return (g.Alpha - 1) / g.Beta
	}
	return math.NaN()
}

// Skewness returns the Pearson's moment coefficient of skewness of the distribution
func (g *Gamma) Skewness() float64 {
	return 2 / math.Sqrt(g.Alpha)
}

// Kurtosis returns the Kurtosis of the distribution
func (g *Gamma) Kurtosis() float64 {
	return 6 / g.Alpha
}

// Entropy returns the Entropy of the distribution
// func (g *Gamma) Entropy() float64 {
// }

// Moment returns the t-th moment of the distribution
func (g *Gamma) Moment(t float64) float64 {
	if t < g.Beta {
		return math.Pow(1-t/g.Beta, -g.Alpha)
	}
	return math.NaN()
}

// Summary returns a string summarising basic info about the distribution
func (g *Gamma) Summary() string {
	dbeg, dend := g.Domain()
	return fmt.Sprintf(`
	X ~ Γ(%f, %f)
		Domain:		[ %f , %f [
		Mean: 		%f
		Var: 		%f
		Skewness: 	%f
		Kurtosis:	%f
`, g.Alpha, g.Beta, dbeg, dend, g.Mean(), g.Var(), g.Skewness(), g.Kurtosis())
}
