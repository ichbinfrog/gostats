package util

import (
	"errors"
)

var (
	ErrBernouilliParam  = errors.New("Invalid parameters, p ∊ [0, 1]")
	ErrBinomialParam    = errors.New("Invalid parameters, p ∊ [0, 1], n > 0")
	ErrExponentialParam = errors.New("Invalid parameters, λ > 0")
	ErrPoissonParam     = errors.New("Invalid parameters, λ > 0")
	ErrGammaParam       = errors.New("Invalid parameters, α > 0, β > 0")
	ErrGeometricParam   = errors.New("Invalid parameters, p ∊ [0, 1]")
	ErrPolyaParam       = errors.New("Invalid parameters, p ∊ [0, 1], r > 0")
	ErrTriangularParam  = errors.New("Invalid parameters, b >= c >= a, b > a")
	ErrUniformParam     = errors.New("Invalid parameters, b > a")
	ErrNormalParam      = errors.New("Invalid parameters, σ^2 > 0")
)
