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
	nextCalendarWeek := nextCalendarWeek()
	for {
		updateIconAndTooltip(<-nextCalendarWeek)
	}
}

func quitOnMenu(){
	mQuit := systray.AddMenuItem("Beenden", "Beendet die Applikation")
	<-mQuit.ClickedCh
	systray.Quit()
}

func updateIconAndTooltip(weekNumber int) {
	systray.SetIcon(generateImage(weekNumber))
	systray.SetTooltip(fmt.Sprintf("Aktuelle Kalenderwoche: %d", weekNumber))
}
