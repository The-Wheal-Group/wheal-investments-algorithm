package funds

//The Fund Parameters type
type FundParameters [4]float64

//The FundsTable type
type FundsTable [3]FundParameters

func GetFunds() FundsTable {
	fund1Parameters := FundParameters{20, 40, 60, 0}
	fund2Parameters := FundParameters{30, 10, 60, 100}
	fund3Parameters := FundParameters{60, 20, 20, 0}
	return FundsTable{fund1Parameters, fund2Parameters, fund3Parameters}
}

func GetDesiredFundParameters() FundParameters {
	return FundParameters{30, 10, 60, 0}
}
