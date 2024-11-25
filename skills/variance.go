package skills

func CalcVariance(data []float64, mean float64) float64 {
	sum := 0.0
	for _, value := range data {
		diff := value - mean
		sum += diff * diff
	}
	return sum / float64(len(data))
}
