package funds

//The Fund Parameters type
type FundParameters [16]float64

//The FundsTable type
type FundsTable [11]FundParameters

func GetFunds() FundsTable {
	return FundsTable{
		FundParameters{100, 0, 0, 0, 0, 0, 24.4, 14, 12.2, 10.7, 17, 8.9, 3.6, 3.5, 3.1, 2.5},
		FundParameters{0, 100, 0, 0, 0, 0, 0, 12.9, 17.4, 4.61, 23, 9.55, 13.61, 4.27, 1.08, 10.55},
		FundParameters{0, 100, 0, 0, 0, 0, 3.3, 3.2, 44.8, 0.7, 22, 20.8, 1.8, 1.7, 0, 0},
		FundParameters{0, 25, 75, 0, 0, 0, 5.7, 15.1, 20.2, 3, 25, 14.8, 6.1, 5, 0, 5.8},
		FundParameters{0, 0, 0, 100, 0, 0, 0, 10, 35, 0, 4.24, 3.6, 0, 0, 12, 6.71},
		FundParameters{0, 0, 0, 0, 100, 0, 6.4, 10, 13, 5.9, 33, 22.9, 0.8, 1.7, 0, 5.5},
		FundParameters{0, 0, 0, 0, 0, 100, 16.84, 4.2, 20.39, 12.24, 21, 5.67, 5.72, 0, 3.39, 6.95},
		FundParameters{100, 0, 0, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 0},
		FundParameters{43.48, 3.49, 8, 1.5, 23.56, 13, 66.12, 3.4, 0, 2.17, 2.6, 24.9, 0, 0, 0, 0},
		FundParameters{47, 0, 19, 13, 0, 12.5, 28.26, 0, 0, 0, 0, 0, 2.28, 45.5, 0, 0},
		FundParameters{100, 0, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
}

func GetDesiredFundParameters() FundParameters {
	return FundParameters{30, 15, 10, 10, 10, 25, 20, 13, 0, 0, 0, 0, 3, 6, 2, 0}
}
