package dist

import (
	"fmt"
	"testing"
)

func TestBernouilli(t *testing.T) {
	dist := &Bernouilli{}
	dist.Init(.3)
	fmt.Println(dist.Summary())

	fmt.Printf("		f(0) = %f", dist.PMF(0))
	fmt.Printf("\n		f(1) = %f\n", dist.PMF(1))
	fmt.Printf("\n		F(-1) = %f", dist.CDF(-1))
	fmt.Printf("\n		F(.3) = %f", dist.CDF(.3))
	fmt.Printf("\n		F(.5) = %f", dist.CDF(.5))
	fmt.Printf("\n		F(1) = %f\n", dist.CDF(1))
	fmt.Printf("\n		Mx(0) = %f", dist.Moment(0))
	fmt.Printf("\n		Mx(1) = %f\n", dist.Moment(1))

	sl := []bool{}
	for i := 0; i < 10; i++ {
		sl = append(sl, dist.Generate())
	}
	fmt.Printf("\n	Generated slice: %v\n\n", sl)
}

func TestBinomial(t *testing.T) {
	testCases := []struct {
		Method string
		N, P   float64
	}{
		{"direct", 24, .3},
		{"direct_poisson", 25, .02},
		{"rejection", 100, .4},
	}
	for _, tc := range testCases {
		dist := &Binomial{}
		t.Run(tc.Method, func(t *testing.T) {
			dist.Init(tc.N, tc.P)
			fmt.Println(dist.Summary())
			fmt.Printf("		f(0) = %f", dist.PMF(0))
			fmt.Printf("\n		f(1) = %f\n", dist.PMF(1))
			fmt.Printf("\n		Mx(0) = %f", dist.Moment(0))
			fmt.Printf("\n		Mx(1) = %f\n", dist.Moment(1))

			sl := []float64{}
			for i := 0; i < 10; i++ {
				sl = append(sl, dist.Generate())
			}
			fmt.Printf("\n	Generated slice: %v\n\n", sl)
		})
	}
}

func TestUniform(t *testing.T) {
	dist := &Uniform{}
	dist.Init(0, 10)
	fmt.Println(dist.Summary())

	fmt.Printf("		f(0) = %f", dist.PMF(0))
	fmt.Printf("\n		f(1) = %f\n", dist.PMF(1))
	fmt.Printf("\n		F(-1) = %f", dist.CDF(-1))
	fmt.Printf("\n		F(.3) = %f", dist.CDF(.3))
	fmt.Printf("\n		F(.5) = %f", dist.CDF(.5))
	fmt.Printf("\n		F(1) = %f\n", dist.CDF(1))
	fmt.Printf("\n		Mx(0) = %f", dist.Moment(0))
	fmt.Printf("\n		Mx(1) = %f\n", dist.Moment(1))

	sl := []float64{}
	for i := 0; i < 10; i++ {
		sl = append(sl, dist.Generate())
	}
	fmt.Printf("\n	Generated slice: %v\n\n", sl)
}

func TestGeometric(t *testing.T) {
	dist := &Geometric{}
	dist.Init(.3)
	fmt.Println(dist.Summary())

	fmt.Printf("		f(0) = %f", dist.PMF(0))
	fmt.Printf("\n		f(1) = %f\n", dist.PMF(1))
	fmt.Printf("\n		F(-1) = %f", dist.CDF(-1))
	fmt.Printf("\n		F(.3) = %f", dist.CDF(.3))
	fmt.Printf("\n		F(.5) = %f", dist.CDF(.5))
	fmt.Printf("\n		F(1) = %f\n", dist.CDF(1))
	fmt.Printf("\n		Mx(0) = %f", dist.Moment(0))
	fmt.Printf("\n		Mx(1) = %f\n", dist.Moment(1))

	sl := []float64{}
	for i := 0; i < 10; i++ {
		sl = append(sl, dist.Generate())
	}
	fmt.Printf("\n	Generated slice: %v\n\n", sl)
}

func TestGamma(t *testing.T) {
	dist := &Gamma{}
	dist.Init(5, 10)
	fmt.Println(dist.Summary())

	fmt.Printf("		f(0) = %f", dist.PMF(0))
	fmt.Printf("\n		f(1) = %f\n", dist.PMF(1))
	fmt.Printf("\n		Mx(0) = %f", dist.Moment(0))
	fmt.Printf("\n		Mx(1) = %f\n", dist.Moment(1))

	sl := []float64{}
	for i := 0; i < 10; i++ {
		sl = append(sl, dist.Generate())
	}
	fmt.Printf("\n	Generated slice: %v\n\n", sl)
}

func TestPoisson(t *testing.T) {
	testCases := []struct {
		Method string
		Lambda float64
	}{
		{"direct", 10},
		{"rejection", 24},
	}

	for _, tc := range testCases {
		t.Run(tc.Method, func(t *testing.T) {
			dist := &Poisson{}
			dist.Init(tc.Lambda)
			fmt.Println(dist.Summary())

			fmt.Printf("		f(0) = %f", dist.PMF(0))
			fmt.Printf("\n		f(1) = %f\n", dist.PMF(1))
			fmt.Printf("\n		F(-1) = %f", dist.CDF(-1))
			fmt.Printf("\n		F(.3) = %f", dist.CDF(.3))
			fmt.Printf("\n		F(.5) = %f", dist.CDF(.5))
			fmt.Printf("\n		F(1) = %f\n", dist.CDF(1))
			fmt.Printf("\n		Mx(0) = %f", dist.Moment(0))
			fmt.Printf("\n		Mx(1) = %f\n", dist.Moment(1))

			sl := []float64{}
			for i := 0; i < 10; i++ {
				sl = append(sl, dist.Generate())
			}
			fmt.Printf("\n	Generated slice: %v\n\n", sl)
		})
	}
}

func TestPolya(t *testing.T) {
	dist := &Polya{}
	dist.Init(5, .3)
	fmt.Println(dist.Summary())

	fmt.Printf("		f(0) = %f", dist.PMF(0))
	fmt.Printf("\n		f(1) = %f\n", dist.PMF(1))
	fmt.Printf("\n		Mx(0) = %f", dist.Moment(0))
	fmt.Printf("\n		Mx(1) = %f\n", dist.Moment(1))

	sl := []float64{}
	for i := 0; i < 10; i++ {
		sl = append(sl, dist.Generate())
	}
	fmt.Printf("\n	Generated slice: %v\n\n", sl)
}
