package ga

import (
	"math/rand"
	"time"
)

type Population struct {
	Chromosomes  []Chromosome
	TotalFitness float64
}

func NewPopulation(size int) Population {
	var population Population

	for index := 0; index < size; index++ {
		population.Chromosomes = append(population.Chromosomes, GenerateRandomChromosome())
		population.Chromosomes[index].Fitness = population.Chromosomes[index].CalculateFitness()
	}

	population.CalculateTotalFitness()

	return population
}

func (population *Population) CalculateTotalFitness() {
	population.TotalFitness = 0

	for _, chromosome := range population.Chromosomes {
		population.TotalFitness += chromosome.CalculateFitness()
	}
}

func (population *Population) CalculateFitness() {
	for index, chromosome := range population.Chromosomes {
		population.Chromosomes[index].Fitness = chromosome.CalculateFitness()
	}
}

func (population *Population) SelectRoulette() Chromosome {

	randomSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randomSource)

	targetRank := random.Float64() * population.TotalFitness

	current := targetRank

	for _, chromosome := range population.Chromosomes {
		current -= chromosome.Fitness

		if current < 0 {
			return chromosome
		}
	}

	return population.Chromosomes[len(population.Chromosomes)-1]
}

func (population *Population) Fittest() Chromosome {
	population.CalculateFitness()

	fittestChromosome := population.Chromosomes[0]

	for _, value := range population.Chromosomes {
		if value.Fitness > fittestChromosome.Fitness {
			fittestChromosome = value
		}
	}

	return fittestChromosome
}
