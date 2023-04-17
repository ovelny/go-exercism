package cars

const carProductionCost = 10_000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (successRate / 100) * float64(productionRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int((successRate / 100) * float64(productionRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	if carsCount >= 10 {
		return uint((carsCount-carsCount%10)*(carProductionCost*0.95) + (carsCount%10)*carProductionCost)
	}
	return uint(carsCount * carProductionCost)
}
