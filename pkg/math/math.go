package math

import (
	"math"
)

func MathRound(x float64) float64 {
	return math.Round(x*100) / 100
}
