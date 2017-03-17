// +build cgo

package cround

import "testing"
import "math"

func TestRound(t *testing.T) {

	tests := []struct {
		x             float64
		toNearestEven float64
		toNearestAway float64
		toZero        float64
		awayFromZero  float64
		toPositiveInf float64
		toNegativeInf float64
	}{
		{+5.5, +6.0, +6.0, +5.0, +6.0, +6.0, +5.0},
		{+2.5, +2.0, +3.0, +2.0, +3.0, +3.0, +2.0},
		{+1.6, +2.0, +2.0, +1.0, +2.0, +2.0, +1.0},
		{+1.1, +1.0, +1.0, +1.0, +2.0, +2.0, +1.0},
		{+1.0, +1.0, +1.0, +1.0, +1.0, +1.0, +1.0},
		{+0.0, +0.0, +0.0, +0.0, +0.0, +0.0, +0.0},
		{-0.0, -0.0, -0.0, -0.0, -0.0, -0.0, -0.0},
		{-1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0},
		{-1.1, -1.0, -1.0, -1.0, -2.0, -1.0, -2.0},
		{-1.6, -2.0, -2.0, -1.0, -2.0, -1.0, -2.0},
		{-2.5, -2.0, -3.0, -2.0, -3.0, -2.0, -3.0},
		{-5.5, -6.0, -6.0, -5.0, -6.0, -5.0, -6.0},
		{math.Inf(+1), math.Inf(+1), math.Inf(+1), math.Inf(+1), math.Inf(+1), math.Inf(+1), math.Inf(+1)},
		{math.Inf(-1), math.Inf(-1), math.Inf(-1), math.Inf(-1), math.Inf(-1), math.Inf(-1), math.Inf(-1)},
	}

	dpTests := []struct {
		x      float64
		dp     float64
		result float64
	}{
		{1234.5678, -4, 0000.0000},
		{1234.5678, -3, 1000.0000},
		{1234.5678, -2, 1200.0000},
		{1234.5678, -1, 1230.0000},
		{1234.5678, +0, 1235.0000},
		{1234.5678, +1, 1234.6000},
		{1234.5678, +2, 1234.5700},
		{1234.5678, +3, 1234.5680},
		{1234.5678, +4, 1234.5678},
	}

	t.Run("Round", func(t *testing.T) {
		for _, tt := range tests {
			result := Round(tt.x)
			if result != tt.toNearestEven {
				t.Errorf("Round(%.1f) = %.1f, expected %.1f", tt.x, result, tt.toNearestEven)
			}
		}
	})

	t.Run("RoundTo", func(t *testing.T) {
		for _, tt := range tests {
			result := RoundTo(tt.x, 0)
			if result != tt.toNearestEven {
				t.Errorf("RoundTo(%.1f, 0) = %.1f, expected %.1f", tt.x, result, tt.toNearestEven)
			}
		}

		for _, tt := range dpTests {
			result := RoundTo(tt.x, tt.dp)
			if result != tt.result {
				t.Errorf("RoundTo(%.4f, %.0f) = %.4f, expected %.4f", tt.x, tt.dp, result, tt.result)
			}
		}
	})

	t.Run("ToNearestEven", func(t *testing.T) {
		for _, tt := range tests {
			result := ToNearestEven(tt.x)
			if result != tt.toNearestEven {
				t.Errorf("ToNearestEven(%.1f) = %.1f, expected %.1f", tt.x, result, tt.toNearestEven)
			}
		}
	})

	t.Run("ToNearestAway", func(t *testing.T) {
		for _, tt := range tests {
			result := ToNearestAway(tt.x)
			if result != tt.toNearestAway {
				t.Errorf("ToNearestAway(%.1f) = %.1f, expected %.1f", tt.x, result, tt.toNearestAway)
			}
		}
	})

	t.Run("ToZero", func(t *testing.T) {
		for _, tt := range tests {
			result := ToZero(tt.x)
			if result != tt.toZero {
				t.Errorf("ToZero(%.1f) = %.1f, expected %.1f", tt.x, result, tt.toZero)
			}
		}
	})

	// There is no AwayFromZero equivalent in C, so it cannot be tested.

	t.Run("ToPositiveInf", func(t *testing.T) {
		for _, tt := range tests {
			result := ToPositiveInf(tt.x)
			if result != tt.toPositiveInf {
				t.Errorf("ToPositiveInf(%.1f) = %.1f, expected %.1f", tt.x, result, tt.toPositiveInf)
			}
		}
	})

	t.Run("ToNegativeInf", func(t *testing.T) {
		for _, tt := range tests {
			result := ToNegativeInf(tt.x)
			if result != tt.toNegativeInf {
				t.Errorf("ToNegativeInf(%.1f) = %.1f, expected %.1f", tt.x, result, tt.toNegativeInf)
			}
		}
	})

}
