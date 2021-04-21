package clock

import "fmt"

type clock struct {
	totalMinutes int
}

const (
	// MinutesInHour : number of minutes in one hour
	MinutesInHour = 60
	// HoursInDay : number of hours in one day
	HoursInDay = 24
	// MinutesInDay : number of minutes in one day
	MinutesInDay = HoursInDay * MinutesInHour
)

// New returns newely initialesed at 00:00
func New(h int, m int) clock {

	return clock{0}.Add(h*MinutesInHour + m)
}

func (c clock) GetMinutes() int {
	return c.totalMinutes % MinutesInHour
}

func (c clock) GetHours() int {
	return c.totalMinutes / MinutesInHour
}

// String returns clock as a string
func (c clock) String() string {

	return fmt.Sprintf("%02d:%02d", c.GetHours(), c.GetMinutes())
}

// Modulos returns the smallest positive module
func Modulos(input int, mod int) int {
	input %= mod
	if input < 0 {
		input += mod
	}

	return input
}

// Add returns a clock with adding minutes from observer clock
func (c clock) Add(m int) clock {

	return clock{Modulos(c.totalMinutes+m, MinutesInDay)}
}

// Subtract returns a clock with substruct minutes from observer clock
func (c clock) Subtract(m int) clock {

	return c.Add(-m)
}
