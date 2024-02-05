package internal

import "time"

func DaysBetween(from, to time.Time) []time.Time {
	if from.After(to) {
		return make([]time.Time, 0)
	}

	days := make([]time.Time, 0)
	for d := toDay(from); d.Before(toDay(to)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days
}

func toDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}
