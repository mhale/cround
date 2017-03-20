// +build cgo

package cround

// #cgo CFLAGS: -frounding-math
// #cgo LDFLAGS: -lm
// #include <fenv.h>
// #include <math.h>
import "C"
import "math"

const Epsilon = 0.0000001

// Round returns the nearest integer value of x, with ties rounded to even integers.
//
// Special cases are:
//	Round(±0) = ±0
//	Round(±Inf) = ±Inf
//	Round(NaN) = NaN
func Round(x float64) float64 {
	return ToNearestEven(x)
}

// RoundTo returns the nearest integer value of x, to dp decimal places, with ties rounded to even integers.
//
// Special cases are:
//	Round(±0, dp) = ±0
//	Round(±Inf, dp) = ±Inf
//	Round(NaN, dp) = NaN
func RoundTo(x float64, dp float64) float64 {
	x = x * math.Pow(10, dp)
	result := ToNearestEven(x)
	return result / math.Pow(10, dp)
}

// ToNearestEven returns the nearest integer value of x, with ties rounded to even integers.
//
// Special cases are:
//	ToNearestEven(±0) = ±0
//	ToNearestEven(±Inf) = ±Inf
//	ToNearestEven(NaN) = NaN
func ToNearestEven(x float64) float64 {
	_, _ = C.fesetround(C.FE_TONEAREST)
	result, _ := C.nearbyint(C.double(x))
	return float64(result)
}

// ToNearestAway returns the nearest integer value of x, with ties rounded away from zero.
//
// Special cases are:
//	ToNearestAway(±0) = ±0
//	ToNearestAway(±Inf) = ±Inf
//	ToNearestAway(NaN) = NaN
func ToNearestAway(x float64) float64 {
	result, _ := C.round(C.double(x))
	return float64(result)
}

// ToZero returns the nearest integer value toward zero. If x is negative it is rounded up, if it is positive it is rounded down.
//
// Special cases are:
//	ToZero(±0) = ±0
//	ToZero(±Inf) = ±Inf
//	ToZero(NaN) = NaN
func ToZero(x float64) float64 {
	_, _ = C.fesetround(C.FE_TOWARDZERO)
	result, _ := C.nearbyint(C.double(x))
	return float64(result)
}

// Note: There is no AwayFromZero equivalent in C.

// ToPositiveInf returns the nearest integer value greater than x.
//
// Special cases are:
//	ToPositiveInf(±0) = ±0
//	ToPositiveInf(±Inf) = ±Inf
//	ToPositiveInf(NaN) = NaN
func ToPositiveInf(x float64) float64 {
	_, _ = C.fesetround(C.FE_UPWARD)
	result, _ := C.nearbyint(C.double(x))
	return float64(result)
}

// ToNegativeInf returns the nearest integer value less than x.
//
// Special cases are:
//	ToNegativeInf(±0) = ±0
//	ToNegativeInf(±Inf) = ±Inf
//	ToNegativeInf(NaN) = NaN
func ToNegativeInf(x float64) float64 {
	_, _ = C.fesetround(C.FE_DOWNWARD)
	result, _ := C.nearbyint(C.double(x))
	return float64(result)
}
