package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	fmt.Println(generateRandomAllocation())

	//score := scoreAllocation(allocation)
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

func generateRandomAllocation() [3]float64 {
	randomSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randomSource)

	var allocation [3]float64

	allocationRemaining := 1.0

	for index := 0; index < 2; index++ {
		randomPercentage := random.Float64()
		randomPercentage = math.Round(randomPercentage*100) / 100
		randomFund := random.Intn(3 - index)

		randomPercentage = math.Round((randomPercentage*allocationRemaining)*100) / 100

		allocationRemaining -= randomPercentage

		runningTotal := -1

		for fundIndex := 0; fundIndex < 3; fundIndex++ {
			if allocation[fundIndex] == 0 {
				runningTotal++
			}

			if runningTotal == randomFund {
				allocation[fundIndex] = randomPercentage
				break
			}
		}
	}

	for fundIndex := 0; fundIndex < 3; fundIndex++ {
		if allocation[fundIndex] == 0 {
			allocation[fundIndex] = math.Round(allocationRemaining*100) / 100
			break
		}
	}

	return allocation
}
