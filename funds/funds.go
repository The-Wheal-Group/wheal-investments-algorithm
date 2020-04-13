package funds

//The Fund Parameters type
type FundParameters [6]float64

//The FundsTable type
type FundsTable [6]FundParameters

func GetFunds() FundsTable {
	fund1Parameters := FundParameters{100, 0, 0, 0, 0, 0}
	fund2Parameters := FundParameters{0, 100, 0, 0, 0, 0}
	fund3Parameters := FundParameters{0, 25, 75, 0, 0, 0}
	fund4Parameters := FundParameters{0, 0, 0, 100, 0, 0}
	fund5Parameters := FundParameters{0, 0, 0, 0, 100, 0}
	fund6Parameters := FundParameters{0, 0, 0, 0, 0, 100}
	return FundsTable{
		fund1Parameters,
		fund2Parameters,
		fund3Parameters,
		fund4Parameters,
		fund5Parameters,
		fund6Parameters}
}

func GetDesiredFundParameters() FundParameters {
	return FundParameters{30, 15, 10, 10, 10, 25}
}
