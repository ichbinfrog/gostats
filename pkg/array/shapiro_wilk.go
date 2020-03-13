package array

import (
	"log"
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

	pi6  = 0.52359877559829887307710723054658381403286156656251763682915743
	c1_0 = .5440e0
	c1_1 = -.39978e0
	c1_2 = .25054e-1
	c1_3 = -.6714e-3

	c2_0 = .13822e1
	c2_1 = -.77857e0
	c2_2 = .62767e-1
	c2_3 = -.20322e-2

	c3_0 = -.15861e1
	c3_1 = -.31082e0
	c3_2 = -.83751e-1
	c3_3 = .38915e-2

	c4_0 = -.4803e0
	c4_1 = -.82676e-1
	c4_2 = .30302e-2

	g1_0 = -.2273e1
	g1_1 = .459e0
)

func swilkFirstPolynomial(n, u float64) float64 {
	return p1_0*math.Pow(u, 5) + p1_1*math.Pow(u, 4) + p1_2*math.Pow(u, 3) + p1_3*math.Pow(u, 2) + p1_4*u + n
}

func swilkSecondPolynomial(n, u float64) float64 {
	return p2_0*math.Pow(u, 5) + p2_1*math.Pow(u, 4) + p2_2*math.Pow(u, 3) + p2_3*math.Pow(u, 2) + p2_4*u + n
}

// ShapiroWilkSignificance implements a translated R version in Golang to compute the shapiro wilk p-value for the given dataset
func ShapiroWilkSignificance(n float64, W float64) float64 {
	if n == 3 {
		if res := pi6 * math.Asin(math.Sqrt(W)-math.Asin(math.Sqrt(3/4))); res > 0 {
			return res
		}
		return math.NaN()
	}

	var m, s float64
	y := math.Log(1 - W)
	xx := math.Log(n)

	d := dist.Normal{}
	d.Init(0, 1)

	if n <= 11.0 {
		gm := g1_0 + g1_1*n
		if y >= gm {
			return math.SmallestNonzeroFloat64
		}
		y = -math.Log(gm - y)
		m = c1_0 + c1_1*n + c1_2*math.Pow(n, 2) + c1_3*math.Pow(n, 3)
		s = math.Exp(c2_0 + c2_1*xx + c2_2*math.Pow(xx, 2) + c2_3*math.Pow(xx, 3))
	} else {
		m = c3_0 + c3_1*xx + c3_2*math.Pow(xx, 2) + c3_3*math.Pow(xx, 3)
		s = math.Exp(c4_0 + c4_1*xx + c4_2*math.Pow(xx, 2))
	}
	return 1 - d.CDF((y-m)/s)
}

// ShapiroWilkStatistic implements AS R94 in Golang to compute the shapiro wilk statistic of a given data array
// TODO: Document + Optimise
func (a *Arrayf64) ShapiroWilkStatistic() float64 {
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
				if n >= 5000 {
					log.Println("[WARN] Sample size too large, Shapiro Wilk statistics might be inaccurate")
				}

				w[n-2] = swilkSecondPolynomial(m[n-2]/sqrtSum, u)
				w[1] = -w[n-2]

				phi = (sum - 2*math.Pow(m[n-1], 2) - 2*math.Pow(m[n-2], 2)) / (1 - (2 * math.Pow(w[n-1], 2)) - (2 * math.Pow(w[n-2], 2)))

				for i := 2; i < n-2; i++ {
					w[i] = m[i] / math.Sqrt(phi)
				}
			}
		}
	}
	W := 0.0
	for i, v := range w {
		W += v * a.Data[i]
	}
	return math.Pow(W, 2) / (a.Var() * float64(n-1))
}
