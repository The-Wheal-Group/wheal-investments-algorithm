package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	"wheal-investments-algorithm/ga"
)

func main() {
	sizeOfPopulation := 100
	population := ga.NewPopulation(sizeOfPopulation)

	fittestEverChromosome := population.Fittest()

	for index := 0; index < 10000; index++ {
		var newPopulation ga.Population

		sort.Slice(population.Chromosomes, func(i, j int) bool { return population.Chromosomes[i].Fitness > population.Chromosomes[j].Fitness })
		elites := ga.Population{
			Chromosomes: append(population.Chromosomes[0:50], fittestEverChromosome),
		}

		for len(newPopulation.Chromosomes) <= sizeOfPopulation {

			chromosome := elites.SelectRoulette()

			randomSource := rand.NewSource(time.Now().UnixNano())
			random := rand.New(randomSource)
			randomNumber := random.Float64()

			if randomNumber < 0.5 {
				mutationRandom := random.Intn(3)

				switch mutationRandom {
				case 1:
					chromosome.MutateIncrement()
				case 2:
					chromosome.MutateSwap()
				case 3:
					chromosome = ga.Crossover(chromosome, population.SelectRoulette())
				}

			}

			if randomNumber > 0.95 {
				chromosome = ga.GenerateRandomChromosome()
			}

			newPopulation.Chromosomes = append(newPopulation.Chromosomes, chromosome)
		}

		newPopulation.CalculateTotalFitness()

		fittest := newPopulation.Fittest()

		if fittest.Fitness > fittestEverChromosome.Fitness {
			fittestEverChromosome = fittest
		}

		fmt.Println(fittest.GetAllocationPercentage(), fittest.Fitness, "Generation:", index)

		population = newPopulation
	}

	fittest := population.Fittest()

	fmt.Println("Answer:", fittest.GetAllocationPercentage())
}
