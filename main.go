package main

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

type AnscombeData struct {
	X1, X2, X3, X4 float64
	Y1, Y2, Y3, Y4 float64
}

// Function to calculate coefficient
// of regression from predicted values
func CalculateCoefficient(regressiondata []stats.Coordinate) (float64, error) {
	// Handle empty input
	if len(regressiondata) == 0 {
		return 0, fmt.Errorf("empty input")
	}

	// Create the variables
	var sumX, sumY, sumXY, sumXX float64
	n := float64(len(regressiondata))

	// Iterate each element in the collection
	// generated by the linear regression function
	for i := range regressiondata {
		x := regressiondata[i].X
		y := regressiondata[i].Y
		sumX += x
		sumY += y
		sumXY += x * y
		sumXX += x * x
	}

	// Calculate the coefficient
	num := (n * sumXY) - (sumX * sumY)
	den := (n * sumXX) - (sumX * sumX)

	coefficient := num / den
	return coefficient, nil
}

func main() {
	// Create the Anscombe Quartet data
	c1 := []stats.Coordinate{
		{X: 10, Y: 8.04},
		{X: 8, Y: 6.95},
		{X: 13, Y: 7.58},
		{X: 9, Y: 8.81},
		{X: 11, Y: 8.33},
		{X: 14, Y: 9.96},
		{X: 6, Y: 7.24},
		{X: 4, Y: 4.26},
		{X: 12, Y: 10.84},
		{X: 7, Y: 4.82},
		{X: 5, Y: 5.68},
	}

	c2 := []stats.Coordinate{
		{X: 10, Y: 9.14},
		{X: 8, Y: 8.14},
		{X: 13, Y: 8.74},
		{X: 9, Y: 8.77},
		{X: 11, Y: 9.26},
		{X: 14, Y: 8.10},
		{X: 6, Y: 6.13},
		{X: 4, Y: 3.10},
		{X: 12, Y: 9.13},
		{X: 7, Y: 7.26},
		{X: 5, Y: 4.74},
	}

	c3 := []stats.Coordinate{
		{X: 10, Y: 7.46},
		{X: 8, Y: 6.77},
		{X: 13, Y: 12.74},
		{X: 9, Y: 7.11},
		{X: 11, Y: 7.81},
		{X: 14, Y: 8.84},
		{X: 6, Y: 6.08},
		{X: 4, Y: 5.39},
		{X: 12, Y: 8.15},
		{X: 7, Y: 6.42},
		{X: 5, Y: 5.73},
	}

	c4 := []stats.Coordinate{
		{X: 8, Y: 6.58},
		{X: 8, Y: 5.76},
		{X: 8, Y: 7.71},
		{X: 8, Y: 8.84},
		{X: 8, Y: 8.47},
		{X: 8, Y: 7.04},
		{X: 8, Y: 5.25},
		{X: 19, Y: 12.50},
		{X: 8, Y: 5.56},
		{X: 8, Y: 7.91},
		{X: 8, Y: 6.89},
	}

	// Fit linear regression models
	r1, _ := stats.LinearRegression(c1)
	coefficient1, _ := CalculateCoefficient(r1)
	fmt.Println("Linear Regression series for c1:", r1, "\nCoefficient of Regression:", coefficient1)

	r2, _ := stats.LinearRegression(c2)
	coefficient2, _ := CalculateCoefficient(r2)
	fmt.Println("Linear Regression series for c2:", r2, "\nCoefficient of Regression:", coefficient2)

	r3, _ := stats.LinearRegression(c3)
	coefficient3, _ := CalculateCoefficient(r3)
	fmt.Println("Linear Regression series for c3:", r3, "\nCoefficient of Regression:", coefficient3)

	r4, _ := stats.LinearRegression(c4)
	coefficient4, _ := CalculateCoefficient(r4)
	fmt.Println("Linear Regression series for c4:", r4, "\nCoefficient of Regression:", coefficient4)
}
