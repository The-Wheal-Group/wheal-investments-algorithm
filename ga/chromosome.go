package ga

import (
	"math"
	"math/rand"
	"time"
)

type Allocation [3]float64

type Chromosome struct {
	Allocation Allocation
	Fitness    float64
}

func (chromosome *Chromosome) CalculateFitness() float64 {
	fund1 := [3]float64{20, 40, 60}
	fund2 := [3]float64{30, 10, 60}
	fund3 := [3]float64{60, 20, 20}
	funds := [3][3]float64{fund1, fund2, fund3}
	desiredAllocation := [3]float64{30, 10, 60}

	var fundsAllocation [3][3]float64

	for regionsIndex := 0; regionsIndex < 3; regionsIndex++ {
		for fundsIndex := 0; fundsIndex < 3; fundsIndex++ {
			fundsAllocation[regionsIndex][fundsIndex] = funds[regionsIndex][fundsIndex] * chromosome.Allocation[regionsIndex]
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

	if difference == 0 {
		difference = 0.00000001
	}

	return 100 / difference
}

func (chromosome *Chromosome) MutateIncrement() {
	randomSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randomSource)

	fundToMutate := random.Intn(3)
	fundToBalance := random.Intn(3)

	plusOrMinus := random.Float64()

	if chromosome.Allocation[fundToMutate] >= 0.01 && chromosome.Allocation[fundToBalance] >= 0.01 &&
		chromosome.Allocation[fundToMutate] <= 0.99 && chromosome.Allocation[fundToBalance] <= 0.90 {
		if plusOrMinus < 0.5 {
			chromosome.Allocation[fundToMutate] += 0.01
			chromosome.Allocation[fundToBalance] -= 0.01
		} else {
			chromosome.Allocation[fundToMutate] -= 0.01
			chromosome.Allocation[fundToBalance] += 0.01
		}
	}
}

func (chromosome *Chromosome) MutateSwap() {
	randomSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randomSource)

	fundToMutate := random.Intn(3)
	fundToSwap := random.Intn(3)

	temp := chromosome.Allocation[fundToMutate]
	chromosome.Allocation[fundToMutate] = chromosome.Allocation[fundToSwap]
	chromosome.Allocation[fundToSwap] = temp
}

func GenerateRandomChromosome() Chromosome {
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

	chromosome := Chromosome{
		Allocation: allocation,
	}

	return chromosome
}
