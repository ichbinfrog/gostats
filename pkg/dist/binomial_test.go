package dist

import (
	"fmt"
	"testing"
)

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
