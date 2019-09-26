package main

import (
	"time"
)

type CalendarWeekIterator struct {
	ChangedCh chan int
}

func currentCalendarWeekIterator() CalendarWeekIterator {
	return calendarWeekIteratorWithCustomProvider(10*time.Minute, func() int {
		_, currentWeekNumber := time.Now().UTC().ISOWeek()
		return currentWeekNumber
	})
}

func calendarWeekIteratorWithCustomProvider(duration time.Duration, currentWeekProvider func() int) CalendarWeekIterator {
	ch := make(chan int)

	go func() {
		var lastWeekNumber = -1
		for ticker := time.NewTicker(duration); ; <-ticker.C {
			currentWeekNumber := currentWeekProvider()
			if lastWeekNumber != currentWeekNumber {
				lastWeekNumber = currentWeekNumber
				ch <- currentWeekNumber
			}
		}
	}()

	return CalendarWeekIterator{ch}
}
