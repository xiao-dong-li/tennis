package game

import (
	"math"
	"time"
)

// CloneMatrix returns a deep copy of the given 2D matrix.
func CloneMatrix(src [][]bool) [][]bool {
	dst := make([][]bool, len(src))
	for i := range src {
		dst[i] = append([]bool(nil), src[i]...)
	}
	return dst
}

// CalcAlpha returns a smoothly oscillating alpha value between 0 and 1.
// It uses a cosine wave to create a fade-in/fade-out effect over a fixed time period.
func CalcAlpha() float64 {
	// Duration of one full fade cycle (2 seconds)
	const period = 2 * time.Second

	// Current position in the fade cycle (0.0 ~ 1.0)
	t := float64(time.Now().UnixNano()%int64(period)) / float64(period)

	// Generate a smooth oscillation curve using cosine:
	// Cosine varies from 1 → -1 → 1, so we map it to 0 → 1 → 0.
	alpha := (math.Cos(t*2*math.Pi) + 1) / 2

	return alpha

	// Convert the normalized alpha (0.0 ~ 1.0) to an alpha value (0 ~ 255)
	//return alpha * 255
}
