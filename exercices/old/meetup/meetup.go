package meetup

import "time"

// WeekSchedule holds week schedule
type WeekSchedule byte

const (
	// First week
	First WeekSchedule = iota
	// Second week
	Second
	// Third week
	Third
	// Fourth week
	Fourth
	// Teenth week
	Teenth
	// Last week
	Last
)

// Day returns day number
func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {

	return 0
}
