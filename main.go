package main

import (
	"fmt"
	"math/rand"
	"time"
	"wheal-investments-algorithm/ga"
)

func main() {
	sizeOfPopulation := 100
	population := ga.NewPopulation(sizeOfPopulation)

	for index := 0; index < 100; index++ {
		var newPopulation ga.Population
		for len(newPopulation.Chromosomes) <= sizeOfPopulation {
			chromosome := population.SelectRoulette()

			randomSource := rand.NewSource(time.Now().UnixNano())
			random := rand.New(randomSource)

			if random.Float64() < 0.4 {
				chromosome.MutateIncrement()
			}

			if random.Float64() > 0.6 {
				chromosome.MutateSwap()
			}

			newPopulation.Chromosomes = append(newPopulation.Chromosomes, chromosome)
		}

		newPopulation.CalculateTotalFitness()

		fmt.Println(newPopulation.TotalFitness)

		population = newPopulation
	}

	fmt.Println("Answer:", population.Fittest())
}
