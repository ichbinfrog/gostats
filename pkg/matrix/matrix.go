package matrix

type Corner struct {
	X0, X1, Y0, Y1 int
}

// Matrixf64 is a bidimensional structure to store float64
type Matrixf64 struct {
	Data       [][]float64
	height     int
	width      int
	Transposed bool
}

// Init allocates the matrix
func (m *Matrixf64) Init(h, w int) {
	m.Data = make([][]float64, h)
	for i := range m.Data {
		m.Data[i] = make([]float64, w)
	}
}

// At returns a pointer to the value at the given index
func (m *Matrixf64) At(i, j int) *float64 {
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
func (m *Matrixf64) T() {
	m.Transposed = !m.Transposed
}

// Height returns the height of the matrix
func (m *Matrixf64) Height() int {
	if m.Transposed {
		return m.width
	}
	return m.height
}

// Width returns the width of the matrix
func (m *Matrixf64) Width() int {
	if m.Transposed {
		return m.height
	}
	return m.width
}

// Add adds two matrix
// Complexity: O(n * m)
//
func Add(a, b *Matrixf64, inplace bool) *Matrixf64 {
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
		return &Matrixf64{
			height:     a.Height(),
			width:      a.Width(),
			Data:       res,
			Transposed: a.Transposed,
		}
	}
	return nil
}

func add(a, b, c *Matrixf64, topA, leftA, topB, leftB, topC, leftC, dim int) {
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			c.Data[topC+i][leftC+j] = a.Data[topA+i][leftA+j] + b.Data[topB+i][leftB+j]
		}
	}
}

func subtract(a, b, c *Matrixf64, topA, leftA, topB, leftB, topC, leftC, dim int) {
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			c.Data[topC+i][leftC+j] = a.Data[topA+i][leftA+j] - b.Data[topB+i][leftB+j]
		}
	}
}

func mul(a, b, c *Matrixf64, topA, leftA, topB, leftB, topC, leftC, dim int) {
	for i := 0; i < dim; i++ {
		for k := 0; k < dim; k++ {
			for j := 0; j < dim; j++ {
				c.Data[topC+i][leftC+j] += a.Data[topA+i][leftA+k] * b.Data[topB+k][leftB+j]
			}
		}
	}
}

func strassen(a, b, c *Matrixf64, topA, leftA, topB, leftB, topC, leftC, dim int) {
	halfDim := dim / 2

	// C12 = A21 - A11
	subtract(a, a, c, topA+halfDim, leftA, topA, leftA, topC, leftC+halfDim, halfDim)
	// C21 = B11 + B12
	add(b, b, c, topB, leftB, topB, leftB+halfDim, topC+halfDim, leftC, halfDim)
	// C22 = C12 * C21
	multiply(c, c, c, topC, leftC+halfDim, topC+halfDim, leftC, topC+halfDim, leftC+halfDim, halfDim)
	//C12 = A12 - A22
	subtract(a, a, c, topA, leftA+halfDim, topA+halfDim, leftA+halfDim, topC, leftC+halfDim, halfDim)
	//C21 = B21 + B22
	add(b, b, c, topB+halfDim, leftB, topB+halfDim, leftB+halfDim, topC+halfDim, leftC, halfDim)
	//C11 = C12 * C21
	multiply(c, c, c, topC, leftC+halfDim, topC+halfDim, leftC, topC, leftC, halfDim)
	//C12 = A11 + A22
	add(a, a, c, topA, leftA, topA+halfDim, leftA+halfDim, topC, leftC+halfDim, halfDim)
	//C21 = B11 + B22
	add(b, b, c, topB, leftB, topB+halfDim, leftB+halfDim, topC+halfDim, leftC, halfDim)

	t1 := &Matrixf64{}
	t1.Init(halfDim, halfDim)

	//T1 = C12*C21
	multiply(c, c, t1, topC, leftC+halfDim, topC+halfDim, leftC, 0, 0, halfDim)
	//C11 = T1 + C11
	add(t1, c, c, 0, 0, topC, leftC, topC, leftC, halfDim)
	//C22 = T1 + C22
	add(t1, c, c, 0, 0, topC+halfDim, leftC+halfDim, topC+halfDim, leftC+halfDim, halfDim)

	t2 := &Matrixf64{}
	t2.Init(halfDim, halfDim)
	//T2 = A21 + A22
	add(a, a, t2, topA+halfDim, leftA, topA+halfDim, leftA+halfDim, 0, 0, halfDim)
	//C21 = T2 * B11
	multiply(t2, b, c, 0, 0, topB, leftB, topC+halfDim, leftC, halfDim)
	//C22 = C22 - C21
	subtract(c, c, c, topC+halfDim, leftC+halfDim, topC+halfDim, leftC, topC+halfDim, leftC+halfDim, halfDim)
	//T1 = B21 - B11
	subtract(b, b, t1, topB+halfDim, leftB, topB, leftB, 0, 0, halfDim)
	//T2 = A22 * T1
	multiply(a, t1, t2, topA+halfDim, leftA+halfDim, 0, 0, 0, 0, halfDim)
	//C21 = C21 + T2
	add(c, t2, c, topC+halfDim, leftC, 0, 0, topC+halfDim, leftC, halfDim)
	//C11 = C11 + T2
	add(c, t2, c, topC, leftC, 0, 0, topC, leftC, halfDim)
	//T1 = B12 - B22
	subtract(b, b, t1, topB, leftB+halfDim, topB+halfDim, leftB+halfDim, 0, 0, halfDim)
	//C12 = A11 * T1
	multiply(a, t1, c, topA, leftA, 0, 0, topC, leftC+halfDim, halfDim)
	//C22 = C22 + C12
	add(c, c, c, topC+halfDim, leftC+halfDim, topC, leftC+halfDim, topC+halfDim, leftC+halfDim, halfDim)
	//T2 = A11 + A12
	add(a, a, t2, topA, leftA, topA, leftA+halfDim, 0, 0, halfDim)
	//T1 = T2 * B22
	multiply(t2, b, t1, 0, 0, topB+halfDim, leftB+halfDim, 0, 0, halfDim)
	//C12 = C12 + T1
	add(c, t1, c, topC, leftC+halfDim, 0, 0, topC, leftC+halfDim, halfDim)
	//C11 = C11 - T1
	subtract(c, t1, c, topC, leftC, 0, 0, topC, leftC, halfDim)
}

func multiply(a, b, c *Matrixf64, topA, leftA, topB, leftB, topC, leftC, dim int) {
	if dim > 2 {
		strassen(a, b, c, topA, leftA, topB, leftB, topC, leftC, dim)
	} else {
		mul(a, b, c, topA, leftA, topB, leftB, topC, leftC, dim)
	}
}

// Mult multiplies two matrix
func Mult(a, b *Matrixf64) *Matrixf64 {
	c := &Matrixf64{}
	c.Init(a.Height(), a.Width())

	// TODO: pretranspose matrix & 0 fill for non square matrices
	strassen(a, b, c, 0, 0, 0, 0, 0, 0, a.Height())
	return c
}
