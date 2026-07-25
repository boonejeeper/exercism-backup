package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100.0)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    var ratePerMinute = float64(productionRate) / 60.0
    return int(ratePerMinute * (successRate / 100.0))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    var buildGroups = carsCount / 10
    var individualCars = carsCount % 10
    return uint(buildGroups * 95000) + uint(individualCars * 10000)
}
