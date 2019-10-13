package main

import (
	"time"

	"github.com/snabb/isoweek"
)

type CalendarWeekIterator struct {
	ChangedCh chan int
}

func currentCalendarWeekIterator() CalendarWeekIterator {
	return calendarWeekIteratorWithCustomProvider(10*time.Minute, currentIsoWeek)
}

func currentIsoWeek() int {
	_, currentWeekNumber := time.Now().ISOWeek()
	return currentWeekNumber
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

func offsetCalendarWeekToDate(idx int) (int, time.Time) {
	return offsetCalendarWeekToDateFromDate(idx, time.Now())
}

func offsetCalendarWeekToDateFromDate(idx int, offset time.Time) (int, time.Time) {
	year, week := offset.ISOWeek()
	for i := 0; i < idx; i++ {
		week++
		if !isoweek.Validate(year, week) {
			week = 1
			year++
		}
	}
	return week, isoweek.StartTime(year, week, time.Local)
}
