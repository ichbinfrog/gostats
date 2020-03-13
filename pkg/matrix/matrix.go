package matrix

// Matrix is a bidimensional structure to store float64
type Matrix struct {
	Data       [][]float64
	height     int
	width      int
	Transposed bool
}

// Init allocates the matrix
func (m *Matrix) Init(h, w int) {
	m.Data = make([][]float64, h)
	for i := range m.Data {
		m.Data[i] = make([]float64, w)
	}
}

// At returns a pointer to the value at the given index
func (m *Matrix) At(i, j int) *float64 {
	if m.Transposed {
		if i >= m.width || j >= m.height {
			return nil
		}
		return &m.Data[j][i]
	}

	if i >= m.height || j >= m.width {
		return nil
	}
	return &m.Data[i][j]
}

// T transposes the given matrix
func (m *Matrix) T() {
	m.Transposed = !m.Transposed
}

// Height returns the height of the matrix
func (m *Matrix) Height() int {
	if m.Transposed {
		return m.width
	}
	return m.height
}

// Width returns the width of the matrix
func (m *Matrix) Width() int {
	if m.Transposed {
		return m.height
	}
	return m.width
}

// Add adds two matrix
// Complexity: O(n * m)
//
func Add(a, b *Matrix, inplace bool) *Matrix {
	if a.Height() == b.Height() && a.Width() == b.width {
		if inplace {
			for i := range a.Data {
				for j := range a.Data[i] {
					a.Data[i][j] += b.Data[i][j]
				}
			}
			return a
		}
		res := make([][]float64, a.Height())
		for i := range res {
			res[i] = make([]float64, b.Width())
			for j := range a.Data[i] {
				res[i][j] = *a.At(i, j) + *b.At(i, j)
			}
		}
		return &Matrix{
			height:     a.Height(),
			width:      a.Width(),
			Data:       res,
			Transposed: a.Transposed,
		}
	}
	return nil
}

// Mult multiples two matrix
// Complexity: O(n^3)
// TODO: Implement Coppersmith/Winograd and or Strassen algorithms
//
func Mult(a, b *Matrix, inplace bool) *Matrix {
	if a.Width() == b.Height() {
		res := make([][]float64, a.Height())
		n := a.Width()
		for i := range res {
			res[i] = make([]float64, b.Width())
			sum := 0.0
			for j := range a.Data[i] {
				for k := 0; k < n; k++ {
					sum += a.Data[i][k] * b.Data[k][j]
				}
				res[i][j] = sum
			}
		}
		if inplace {
			a.Data = res
			return a
		}
		return &Matrix{
			height:     a.Height(),
			width:      b.Width(),
			Data:       res,
			Transposed: false,
		}
	}
	return nil
}
