package array

import (
	"math"

	"github.com/ichbinfrog/statistics/pkg/dist"
)

const (
	p1_0 = -2.706056
	p1_1 = 4.434685
	p1_2 = -2.071190
	p1_3 = -0.147981
	p1_4 = 0.221157

	p2_0 = -3.582633
	p2_1 = 5.682633
	p2_2 = -1.752461
	p2_3 = -0.293762
	p2_4 = 0.042981

	pi6 = 0.52359877559829887307710723054658381403286156656251763682915743
)

func swilkFirstPolynomial(n, u float64) float64 {
	return p1_0*math.Pow(u, 5) + p1_1*math.Pow(u, 4) + p1_2*math.Pow(u, 3) + p1_3*math.Pow(u, 2) + p1_4*u + n
}

func swilkSecondPolynomial(n, u float64) float64 {
	return p2_0*math.Pow(u, 5) + p2_1*math.Pow(u, 4) + p2_2*math.Pow(u, 3) + p2_3*math.Pow(u, 2) + p2_4*u + n
}

// ShapiroWilk implements AS R94 in Golang
// TODO: Document + Optimise
func (a *Arrayf64) ShapiroWilk() float64 {
	// ROYSTON, Patrick. Remark AS R94: A remark on algorithm AS 181: The W-test for normality. Journal of the Royal Statistical Society. Series C (Applied Statistics), 1995, vol. 44, no 4, p. 547-551.
	n := int(a.Length)
	m := make([]float64, n)
	d := dist.Normal{}
	d.Init(0, 1)

	if a.Length < 3 {
		return math.NaN()
	}

	sum := 0.0
	for i := 0; i < n; i++ {
		m[i] = d.Quantile(((float64(i + 1)) - (3.0 / 8.0)) / (float64(n) + .25))
		sum += math.Pow(m[i], 2)
	}
	sqrtSum := math.Sqrt(sum)

	// Normalise by square root
	w := make([]float64, n)
	k := a.Kurtosis()
	for i := 0; i < int(a.Length); i++ {
		w[i] = m[i] / sqrtSum
	}

	W := 0.0
	if k > 3.0 {
		// Shapiro-Francia test for leptokurtic samples
	} else {
		// Shapiro-Wilk test for platykurtic samples
		if n == 3 {
			w[0] = math.Sqrt(.5)
			w[2] = -w[0]
		} else {
			var phi, u float64
			u = 1 / math.Sqrt(float64(n))
			w[n-1] = swilkFirstPolynomial(m[n-1]/sqrtSum, u)
			w[0] = -w[n-1]

			if n <= 5 {
				// N == 4 || N == 5
				phi = (sum - 2*math.Pow(m[n-1], 2)) / (1 - (2 * math.Pow(w[n-1], 2)))
				for i := 1; i < n-1; i++ {
					w[i] = m[i] / math.Sqrt(phi)
				}
			} else {
				// N >= 6
				w[n-2] = swilkSecondPolynomial(m[n-2]/sqrtSum, u)
				w[1] = -w[n-2]

				phi = (sum - 2*math.Pow(m[n-1], 2) - 2*math.Pow(m[n-2], 2)) / (1 - (2 * math.Pow(w[n-1], 2)) - (2 * math.Pow(w[n-2], 2)))

				for i := 2; i < n-2; i++ {
					w[i] = m[i] / math.Sqrt(phi)
				}
			}
		}
	}
	for i, v := range w {
		W += v * a.Data[i]
	}
	return math.Pow(W, 2) / (a.Var() * float64(n-1))
}
