package ga

//The Populatio type
type Population struct {
	Chromosomes  []Chromosome
	TotalFitness float64
}

//Create a new population of a given size
func NewPopulation(size int) Population {
	var population Population

	//Loop through the new population
	for index := 0; index < size; index++ {
		//Generate a new chromosome and add to the population
		population.Chromosomes = append(population.Chromosomes, GenerateRandomChromosome())
	}

	//Calculate the total fitness of the new population
	population.CalculateFitness()

	//Return the new population
	return population
}

//Calculate the total fitness of the population
func (population *Population) CalculateFitness() {
	//Initialise the population total fitness
	population.TotalFitness = 0

	//Loop through the population
	for index, chromosome := range population.Chromosomes {
		//Calculate the fitness of the chromosome
		chromosomeFitness := chromosome.CalculateFitness()

		//Set the chromosome fitness
		population.Chromosomes[index].Fitness = chromosomeFitness

		//Add the individual chromosome fitness to the population total fitness
		population.TotalFitness += chromosome.Fitness
	}
}

//Select a random chromosome from the population based on fitness
func (population *Population) SelectRoulette() Chromosome {

	//Get a target rank in the population
	targetRank := Random().Float64() * population.TotalFitness

	//Set the currennt rank equal to the target rank
	currentRank := targetRank

	//Loop through the population of chromosomes
	for _, chromosome := range population.Chromosomes {
		//Get closer to the target rank
		currentRank -= chromosome.Fitness

		//If met the target rank
		if currentRank < 0 {
			return chromosome
		}
	}

	//If target rank cannot be met, return last chromosome in population
	return population.Chromosomes[len(population.Chromosomes)-1]
}

//Get the fittest chromosome
func (population *Population) Fittest() Chromosome {
	//Initialise the fittest chromosome
	fittestChromosome := population.Chromosomes[0]

	//If the fitness of the population hasn't been calculated (should never happen)
	if fittestChromosome.Fitness == 0.0 {
		population.CalculateFitness()
	}

	//Loop through the population
	for _, chromosome := range population.Chromosomes {
		//If the chromosome is fitter than the current fittest
		if chromosome.Fitness > fittestChromosome.Fitness {
			//Set the fittest chromosome to be this chromosome
			fittestChromosome = chromosome
		}
	}

	//Retrun the fittest chromosome
	return fittestChromosome
}
