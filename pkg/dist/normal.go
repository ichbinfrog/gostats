package dist

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/ichbinfrog/statistics/pkg/util"
)

// Normal represents the Poisson distribution
// Continuous probability distribution function as follows:
// 		X ~ N(μ	, σ)
//
type Normal struct {
	Mu, Sigma float64
}

// Init intialises a Bernoulli distribution
func (n *Normal) Init(mu, sigma float64) error {
	if sigma <= 0 {
		return util.ErrNormalParam
	}
	n.Mu, n.Sigma = mu, sigma
	return nil
}

// Generate creates one sample of the Poisson distribution
func (n *Normal) Generate() float64 {
	return rand.NormFloat64()*n.Sigma + n.Mu
}

// Domain returns the definition domain of the distribution
func (n *Normal) Domain() (float64, float64) {
	return math.Inf(-1), math.Inf(0)
}

// PMF returns the probability mass function value of a given k
func (n *Normal) PMF(x float64) float64 {
	return math.Exp(-math.Pow(((x-n.Mu)/n.Sigma), 2) / 2)
}

// CDF returns the Cumulative distribution function value of a given k
func (n *Normal) CDF(x float64) float64 {
	return (.5 + .5*math.Erf((x-n.Mu)/(n.Sigma*math.Sqrt(2))))
}

// Mean returns the mean of the distribution
func (n *Normal) Mean() float64 {
	return n.Mu
}

// Quantile returns the pth-quantile of the distribution
func (n *Normal) Quantile(p float64) float64 {
	return n.Mu + n.Sigma*math.Sqrt(2)*math.Erfinv(2*p-1)
}

// Median returns the median of the distribution
func (n *Normal) Median() float64 {
	return n.Mu
}

// Var returns the variance of the distribution
func (n *Normal) Var() float64 {
	return math.Sqrt(n.Sigma)
}

// Skewness returns the Pearson's moment coefficient of skewness of the distribution
func (n *Normal) Skewness() float64 {
	return 0
}

// Kurtosis returns the Kurtosis of the distribution
func (n *Normal) Kurtosis() float64 {
	return 0
}

// Entropy returns the Entropy of the distribution
func (n *Normal) Entropy() float64 {
	return (math.Log(2 * math.Pi * math.Exp(1) * math.Pow(n.Sigma, 2)))
}

// Moment returns the t-th moment of the distribution
func (n *Normal) Moment(t float64) float64 {
	return math.Exp(n.Mu*t + math.Pow(n.Sigma, 2)*math.Pow(t, 2)/2)
}

// FisherI returns the Fisher Information of the distribution
func (n *Normal) FisherI() [][]float64 {
	return [][]float64{
		[]float64{1 / math.Pow(n.Sigma, 2), 0},
		[]float64{0, 2 / math.Pow(n.Sigma, 2)},
	}
}

// Summary returns a string summarising basic info about the distribution
func (n *Normal) Summary() string {
	dbeg, dend := n.Domain()
	return fmt.Sprintf(`
	X ~ N(%f, %f)
		Domain:			] %f , %f [
		Mean: 			%f
		Median:			%f
		Var: 			%f
		Skewness: 		%f
		Kurtosis:		%f
		FisherInfo:		%v
`, n.Mu, n.Sigma, dbeg, dend, n.Mean(), n.Median(), n.Var(), n.Skewness(), n.Kurtosis(), n.FisherI())
}
