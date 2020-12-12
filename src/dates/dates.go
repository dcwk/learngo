package dates

import "math"

const DaysInWeek = 7

func WeekDoDays(weeks int) int {
	return weeks * DaysInWeek
}

func DaysToWeeks(days int) float64 {
	return math.Ceil(float64(days) / DaysInWeek)
}
