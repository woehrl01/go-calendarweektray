package main

import (
	"fmt"
	"time"

	"github.com/getlantern/systray"
	"github.com/goodsign/monday"
	"github.com/snabb/isoweek"
)

func main() {
	EnableDpiAwareness()

	systray.Run(onReady, func() {})
}

var menus []*systray.MenuItem

func onReady() {
	systray.SetTitle("Kalenderwoche")

	addMenuItemsForUpcomingCalendarWeekDates()

	go keepWeekNumberIconUpToDate()
	go quitOnMenu()
}

func addMenuItemsForUpcomingCalendarWeekDates() {
	for i := 0; i < 15; i++ {
		menus = append(menus, systray.AddMenuItem("refresh", ""))
	}

	systray.AddSeparator()
}

func keepWeekNumberIconUpToDate() {
	calendarWeek := currentCalendarWeekIterator()
	for {
		updateIconAndTooltip(<-calendarWeek.ChangedCh)
	}
}

func quitOnMenu() {
	quitMenuItem := systray.AddMenuItem("Beenden", "Beendet die Applikation")
	<-quitMenuItem.ClickedCh
	systray.Quit()
}

func updateIconAndTooltip(weekNumber int) {
	systray.SetIcon(generateImage(weekNumber))
	systray.SetTooltip(fmt.Sprintf("Aktuelle Kalenderwoche: %d", weekNumber))

	refreshUpcomingCalendarWeekItems()
}

func refreshUpcomingCalendarWeekItems() {
	year, week := time.Now().ISOWeek()

	for index, _ := range menus {
		week++
		startDate := isoweek.StartTime(year, week, time.Local)

		if !isoweek.Validate(year, week) {
			week = 1
			year++
		}

		text := fmt.Sprintf("KW %d: %s", week, monday.Format(startDate, "02. January 2006", monday.LocaleDeDE))
		menus[index].SetTitle(text)
	}
}
