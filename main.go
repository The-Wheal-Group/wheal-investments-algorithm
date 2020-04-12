package main

import (
	"fmt"
	"math"
)

func main() {
	allocation := [3]float64{0.0, 0.9, 0.1}

	score := scoreAllocation(allocation)

	fmt.Println(score)
}

func scoreAllocation(allocation [3]float64) float64 {
	fund1 := [3]float64{20, 40, 60}
	fund2 := [3]float64{30, 10, 60}
	fund3 := [3]float64{60, 20, 20}
	funds := [3][3]float64{fund1, fund2, fund3}
	desiredAllocation := [3]float64{30, 10, 60}

	var fundsAllocation [3][3]float64

	for regionsIndex := 0; regionsIndex < 3; regionsIndex++ {
		for fundsIndex := 0; fundsIndex < 3; fundsIndex++ {
			fundsAllocation[regionsIndex][fundsIndex] = funds[regionsIndex][fundsIndex] * allocation[regionsIndex]
		}
	}

	var actualAllocation [3]float64

	for regionsIndex := 0; regionsIndex < 3; regionsIndex++ {
		for fundsIndex := 0; fundsIndex < 3; fundsIndex++ {
			actualAllocation[regionsIndex] += fundsAllocation[fundsIndex][regionsIndex]
		}
	}

	difference := 0.0

	for index := 0; index < 3; index++ {
		difference += math.Abs(desiredAllocation[index] - actualAllocation[index])
	}

	return difference
}
