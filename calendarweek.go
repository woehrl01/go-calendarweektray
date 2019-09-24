package main

import (
	"time"
)

type CalendarWeekIterator struct{
	ChangedCh chan int
}

func currentCalendarWeekIterator() CalendarWeekIterator {
	const duration = 10 * time.Minute;
	
	ch := make(chan int)

	go func() {
		var lastWeekNumber = -1;
		for ticker := time.NewTicker(duration) ; ; <-ticker.C {
			_, currentWeekNumber := time.Now().UTC().ISOWeek()
			if(lastWeekNumber != currentWeekNumber){
				lastWeekNumber = currentWeekNumber
				ch <- currentWeekNumber
			}
		}
	}()

	return CalendarWeekIterator{ch};
}
