package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var totalBirds int
	for i := 0; i < len(birdsPerDay); i++ {
        totalBirds += birdsPerDay[i]
    }
    return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	firstIndex := (week * 7) - 7
    lastIndex := (week * 7) -1
	var totalBirds int
    for i := firstIndex; i <= lastIndex; i++ {
        totalBirds += birdsPerDay[i]
    } 
    return totalBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
        if i % 2 == 0 {
            birdsPerDay[i]++
        }
    }
    return birdsPerDay
}
