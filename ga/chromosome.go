package ga

import (
	"math"
	"wheal-investments-algorithm/funds"
)

//The FundAllocation type
type FundAllocation [3]float64

//The Chromosome type
type Chromosome struct {
	FundAllocation FundAllocation
	Fitness        float64
}

//Calculate the fitness of a chromosome
func (chromosome *Chromosome) CalculateFitness() float64 {
	//Get the funds table
	fundsTable := funds.GetFunds()

	//Get the desired fund parameters
	desiredFundParameters := funds.GetDesiredFundParameters()

	//A table used to store the actual allocation values for each fund
	var allocatedFundsTable funds.FundsTable

	//Get the chromosome percentage fund allocation
	percentageFundAllocation := chromosome.GetFundAllocationPercentage()

	//Loop through all the funds
	for fundIndex := 0; fundIndex < len(fundsTable); fundIndex++ {
		//Loop through all the parameters
		for parameterIndex := 0; parameterIndex < len(desiredFundParameters); parameterIndex++ {
			//Calculate the actual allocation values for each fund
			allocatedFundsTable[fundIndex][parameterIndex] = fundsTable[fundIndex][parameterIndex] * percentageFundAllocation[fundIndex]
		}
	}

	//Used to store the fund parameters of the actual fund created
	var actualFundParameters funds.FundParameters

	//Loop through all the parameters
	for parameterIndex := 0; parameterIndex < len(desiredFundParameters); parameterIndex++ {
		//Loop through all the funds
		for fundIndex := 0; fundIndex < len(fundsTable); fundIndex++ {
			actualFundParameters[parameterIndex] += allocatedFundsTable[fundIndex][parameterIndex]
		}
	}

	//Store the difference between the actual and desired parameters
	difference := 0.0

	//Loop through all the parameters
	for parameterIndex := 0; parameterIndex < len(desiredFundParameters); parameterIndex++ {
		//Ignore if the desired parameter equals zero
		if desiredFundParameters[parameterIndex] != 0 {
			//Calculate the difference between the actual and desired parameters and make positive
			difference += math.Abs(desiredFundParameters[parameterIndex] - actualFundParameters[parameterIndex])
		}
	}

	//Avoid a divide by zero bug (i.e. fitness of infinity)
	if difference == 0 {
		difference = 0.00000001
	}

	//Return the fitness (the bigger the better)
	return 100 / difference
}

//Get the weighted fund allocation (won't add up to 100% without this function)
func (chromosome *Chromosome) GetFundAllocationPercentage() FundAllocation {
	//Create the percentage allocation
	var percentageAllocation FundAllocation

	//Initialise the total allocation
	total := 0.0

	//Loop through all the funds and calculate the total allocation
	for fundIndex := 0; fundIndex < len(chromosome.FundAllocation); fundIndex++ {
		total += chromosome.FundAllocation[fundIndex]
	}

	//Loop through all the funds and calculate the weighted allocation (to equal 1)
	for index := 0; index < len(chromosome.FundAllocation); index++ {
		percentageAllocation[index] = chromosome.FundAllocation[index] / total
	}

	//Return the percenntage allocation
	return percentageAllocation
}

//Generate a random chromosome
func GenerateRandomChromosome() Chromosome {

	//The new fund allocation for the chromosome
	var fundAllocation FundAllocation

	//Loop through all the funds
	for index := 0; index < len(fundAllocation); index++ {
		//Give a fund a ranndom allocation
		fundAllocation[index] = Random().Float64()
	}

	//Create the new chromosome
	chromosome := Chromosome{
		FundAllocation: fundAllocation,
	}

	//Return the new chromomosome
	return chromosome
}

//Mutate the chromosome by incrementing a random value
func (chromosome *Chromosome) MutateIncrement() {
	//Select a random fund to mutate
	fundToMutate := Random().Intn(len(chromosome.FundAllocation) - 1)

	//Select a random fund to balance
	fundToBalance := Random().Intn(len(chromosome.FundAllocation) - 1)

	//If the mutation won't cause the fund allocation to go under zero
	if chromosome.FundAllocation[fundToMutate] >= 0.01 && chromosome.FundAllocation[fundToBalance] >= 0.01 {
		//If the mutation won't cause the fund allocation to go over one
		if chromosome.FundAllocation[fundToMutate] <= 0.99 && chromosome.FundAllocation[fundToBalance] <= 0.99 {
			//Mutate the fund allocations
			chromosome.FundAllocation[fundToMutate] += 0.01
			chromosome.FundAllocation[fundToBalance] -= 0.01
		}
	}
}

//Mutate the chromosome by swapping a ranndom value
func (chromosome *Chromosome) MutateSwap() {
	//Select a random fund to mutate
	fundToMutate := Random().Intn(len(chromosome.FundAllocation) - 1)

	//Select a random fund to mutate
	fundToSwap := Random().Intn(len(chromosome.FundAllocation) - 1)

	//Swap the fund allocations
	temp := chromosome.FundAllocation[fundToMutate]
	chromosome.FundAllocation[fundToMutate] = chromosome.FundAllocation[fundToSwap]
	chromosome.FundAllocation[fundToSwap] = temp
}

//Single crossover two chromosomes
func SingleCrossover(parent1 Chromosome, parent2 Chromosome) Chromosome {

	//Randomly select a fund to crossover
	fundToCrossover := Random().Intn(len(parent1.FundAllocation) - 1)

	//Crossover the fund allocation at the random point
	child := parent1
	child.FundAllocation[fundToCrossover] = parent2.FundAllocation[fundToCrossover]

	//Return the child chromosome
	return child
}

//Mutliple crossover two chromosomes
func MultipleCrossover(parent1 Chromosome, parent2 Chromosome) Chromosome {

	//Randomly select a crossover point
	crossoverPoint := Random().Intn(len(parent1.FundAllocation) - 1)

	//The child equals the first parent
	child := parent1

	//Loop through parent 2 allocations
	for index, value := range parent2.FundAllocation {
		//If the index is larger than the crossover point
		if index >= crossoverPoint {
			//Set the child allocation equal to parent allocation
			child.FundAllocation[index] = value
		}
	}

	//Return the child chromosome
	return child
}
