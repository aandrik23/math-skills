package skills

import (
	"math"
	"sort"
)

func CalcMedian(data []float64) float64 {
	sort.Float64s(data)

	n := len(data) / 2
	if len(data)%2 == 0 {
		return math.Round(float64(data[n-1])+float64(data[n])) / 2.0
	} else {
		return data[n]
	}
}
