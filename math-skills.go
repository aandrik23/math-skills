package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"math-skills/skills"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run your-program.go <file-path>")
		return
	}

	filePath := os.Args[1]
	data, err := readFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	average := int(math.Round(skills.CalcAverage(data)))
	median := int(math.Round(skills.CalcMedian(data)))
	variance := int(math.Round(skills.CalcVariance(data, float64(average))))

	fmt.Println("Average:", average)
	fmt.Println("Median:", median)
	fmt.Println("Variance:", variance)
	fmt.Println("Standard Deviation:", int(math.Round(math.Sqrt(float64(variance)))))
}

func readFile(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	var data []float64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number in file: %s", line)
		}
		data = append(data, value)

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
