package booking

import "time"

// Generic function to parse various date formats along the exercise.
func dateParser(layout, date string) time.Time {
	time, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return time
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	return dateParser("1/2/2006 15:04:05", date)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t := dateParser("January 2, 2006 15:04:05", date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t := dateParser("Monday, January 2, 2006 15:04:05", date)
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t := dateParser("1/2/2006 15:04:05", date)
	return t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
