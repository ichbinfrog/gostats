package dist

import (
	"fmt"
	"testing"
)

func TestTriangular(t *testing.T) {
	dist := &Triangular{}
	dist.Init(1, 3, 2)
	fmt.Println(dist.Summary())

	fmt.Printf("		f(0) = %f", dist.PMF(0))
	fmt.Printf("\n		f(1) = %f\n", dist.PMF(1))
	fmt.Printf("\n		F(0) = %f", dist.CDF(0))
	fmt.Printf("\n		F(1) = %f", dist.CDF(1))
	fmt.Printf("\n		F(2) = %f", dist.CDF(2))
	fmt.Printf("\n		F(3) = %f\n", dist.CDF(3))
	fmt.Printf("\n		Mx(0) = %f", dist.Moment(0))
	fmt.Printf("\n		Mx(1) = %f\n", dist.Moment(1))

	sl := []float64{}
	for i := 0; i < 10; i++ {
		sl = append(sl, dist.Generate())
	}
	fmt.Printf("\n	Generated slice: %v\n\n", sl)
}
