package ga

import (
	"math"
)

//The FundAllocation type
type FundAllocation [3]float64

//The Fund Parameters type
type FundParameters [3]float64

//The FundsTable type
type FundsTable [3]FundParameters

//The Chromosome type
type Chromosome struct {
	FundAllocation FundAllocation
	Fitness        float64
}

//Calculate the fitness of a chromosome
func (chromosome *Chromosome) CalculateFitness() float64 {
	fund1Parameters := FundParameters{20, 40, 60}
	fund2Parameters := FundParameters{30, 10, 60}
	fund3Parameters := FundParameters{60, 20, 20}
	funds := FundsTable{fund1Parameters, fund2Parameters, fund3Parameters}
	desiredFundParameters := FundParameters{30, 10, 60}

	//A table used to store the actual allocation values for each fund
	var allocatedFundsTable FundsTable

	//Get the chromosome percentage fund allocation
	percentageFundAllocation := chromosome.GetFundAllocationPercentage()

	//Loop through all the funds
	for fundIndex := 0; fundIndex < len(funds); fundIndex++ {
		//Loop through all the parameters
		for parameterIndex := 0; parameterIndex < len(desiredFundParameters); parameterIndex++ {
			//Calculate the actual allocation values for each fund
			allocatedFundsTable[fundIndex][parameterIndex] = funds[fundIndex][parameterIndex] * percentageFundAllocation[fundIndex]
		}
	}

	//Used to store the fund parameters of the actual fund created
	var actualFundParameters FundParameters

	//Loop through all the parameters
	for parameterIndex := 0; parameterIndex < len(desiredFundParameters); parameterIndex++ {
		//Loop through all the funds
		for fundIndex := 0; fundIndex < len(funds); fundIndex++ {
			actualFundParameters[parameterIndex] += allocatedFundsTable[fundIndex][parameterIndex]
		}
	}

	//Store the difference between the actual and desired parameters
	difference := 0.0

	//Loop through all the parameters
	for parameterIndex := 0; parameterIndex < len(desiredFundParameters); parameterIndex++ {
		//Calculate the difference between the actual and desired parameters and make positivre
		difference += math.Abs(desiredFundParameters[parameterIndex] - actualFundParameters[parameterIndex])
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
