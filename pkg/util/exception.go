package util

import (
	"errors"
)

var (
	// ErrBernoulliParam is returned when the probability parameter is not within the [0, 1] range for the Bernoulli Distribution to be initialized
	ErrBernoulliParam = errors.New("Invalid parameters, p ∊ [0, 1]")
	// ErrBinomialParam is returned when the probability parameter is not within the [0, 1] range for the Binomial Distribution to be initialized
	ErrBinomialParam = errors.New("Invalid parameters, p ∊ [0, 1], n > 0")

	// ErrExponentialParam is returned when the λ parameter is not greater than 0 for the Exponential Distribution to be initialized
	ErrExponentialParam = errors.New("Invalid parameters, λ > 0")

	// ErrPoissonParam is returned when the λ parameter is not greater than 0 for the Poisson Distribution to be initialized
	ErrPoissonParam = errors.New("Invalid parameters, λ > 0")

	// ErrGammaParam is returned when the α and β parameter are not greater than 0 for the Gamma Distribution to be initialized
	ErrGammaParam = errors.New("Invalid parameters, α > 0, β > 0")

	// ErrGeometricParam is returned when the probability parameter is not within the [0, 1] range for the Geometric distribution to be initialized
	ErrGeometricParam = errors.New("Invalid parameters, p ∊ [0, 1]")

	// ErrPolyaParam is returned when the probability parameter is not within the [0, 1] range and r is not greater than 0 for the Polya (negative binomial) distribution to be initialized
	ErrPolyaParam = errors.New("Invalid parameters, p ∊ [0, 1], r > 0")

	// ErrTriangularParam is returned when the three points do not follow the b >= c >= a (with b != a) for the Triangular distribution to be initialized
	ErrTriangularParam = errors.New("Invalid parameters, b >= c >= a, b > a")

	// ErrUniformParam is returned when the upper limit is not strictly higher than the lower for the Uniform distribution to be initialized
	ErrUniformParam = errors.New("Invalid parameters, b > a")

	// ErrNormalParam is returned when the variance is not greater than 0 for the Normal distribution to be initialized
	ErrNormalParam = errors.New("Invalid parameters, σ^2 > 0")
)
