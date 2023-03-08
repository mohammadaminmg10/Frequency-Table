package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	data := []float64{
		128, 130, 130, 133, 134, 135, 136,
		136, 138, 138, 138, 141, 141, 141,
		142, 142, 142, 143, 143, 143, 143,
		143, 144, 144, 145, 145, 145, 145,
		145, 145, 129,
	}
	numberOfData := len(data)
	sort.Float64s(data)

	minVal := data[0]
	maxVal := data[numberOfData-1]

	Range := maxVal - minVal + 1
	levelNum := 1 + int(math.Ceil(3.3*math.Log10(float64(numberOfData))))
	height := math.Ceil(Range / float64(levelNum))

	var table [][]float64
	for i := 1; i <= levelNum; i++ {
		lowerBound := minVal + float64(i-1)*height
		upperBound := lowerBound + height - 1
		classLimit := []float64{lowerBound, upperBound}
		classBoundary := []float64{lowerBound - 0.5, upperBound + 0.5}
		classData := []float64{}
		for _, val := range data {
			if val >= lowerBound && val <= upperBound {
				classData = append(classData, val)
			}
		}
		classMark := 0.0
		if len(classData) > 0 {
			sum := 0.0
			for _, val := range classData {
				sum += val
			}
			classMark = sum / float64(len(classData))
		}
		classFreq := float64(len(classData))
		row := []float64{classLimit[0], classLimit[1], classBoundary[0], classBoundary[1], classMark, classFreq}
		table = append(table, row)
	}

	fmt.Println("Class Limit\t\tClass Boundary\t\tClass Mark\tFrequency")
	for _, row := range table {
		fmt.Printf("%.1f-%.1f\t\t%.1f-%.1f\t\t%.2f\t\t%.0f\t\n", row[0], row[1], row[2], row[3], row[4], row[5])
	}
}
