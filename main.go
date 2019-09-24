package main

import (
	"fmt"
	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, func(){})
}

func onReady() {
	systray.SetTitle("Kalenderwoche")

	go keepWeekNumberIconUpToDate()
	go quitOnMenu()
}

func keepWeekNumberIconUpToDate(){
	calendarWeek := currentCalendarWeekIterator()
	for {
		updateIconAndTooltip(<-calendarWeek.ChangedCh)
	}
}

func quitOnMenu(){
	quitMenuItem := systray.AddMenuItem("Beenden", "Beendet die Applikation")
	<-quitMenuItem.ClickedCh
	systray.Quit()
}

func updateIconAndTooltip(weekNumber int) {
	systray.SetIcon(generateImage(weekNumber))
	systray.SetTooltip(fmt.Sprintf("Aktuelle Kalenderwoche: %d", weekNumber))
}
