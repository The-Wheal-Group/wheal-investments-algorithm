package main

func main() {
	//fund1 := [3]int{20, 40, 60}
	//fund2 := [3]int{30, 10, 60}
	desired := [3]int{30, 10, 60}

	allocation := [3]int{30, 20, 50}

	scoreAllocation(allocation, desired)
}

func scoreAllocation(allocation [3]int, desired [3]int) {

	difference := 0

	for index, value := range desired {
		if value > allocation[index] {
			difference += value - allocation[index]
		} else {
			difference += allocation[index] - value
		}
	}
}
