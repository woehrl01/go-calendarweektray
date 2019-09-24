package main

import (
	"time"
)

func nextCalendarWeek() chan int{
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

	return ch;
}
