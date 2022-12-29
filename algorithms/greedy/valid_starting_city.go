package greedy

func ValidStartingCity(distances []int, fuel []int, mpg int) int {
	numberOfCities := len(distances)
	milesRemaining := 0

	indexOfStartingCityCandidate := 0
	milesRemainingAtStartingCityCandidate := 0

	for cityIdx := 1; cityIdx < numberOfCities; cityIdx++ {
		distanceFromPreviousCity := distances[cityIdx-1]
		fuelFromPreviousCity := fuel[cityIdx-1]
		milesRemaining += fuelFromPreviousCity*mpg - distanceFromPreviousCity

		if milesRemaining < milesRemainingAtStartingCityCandidate {
			milesRemainingAtStartingCityCandidate = milesRemaining
			indexOfStartingCityCandidate = cityIdx
		}
	}

	return indexOfStartingCityCandidate
}
