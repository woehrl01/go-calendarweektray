package main

import (
	"fmt"

	"github.com/getlantern/systray"
	"github.com/goodsign/monday"
)

var (
	sha1ver   string // sha1 revision used to build the program
	buildTime string // when the executable was built
	semVer    string // the version of the build
)

func main() {
	enableDpiAwareness()

	systray.Run(onReady, func() {})
}

var menus []*systray.MenuItem

func onReady() {
	systray.SetTitle("Kalenderwoche")

	const numberOfEntries = 15
	addMenuItemsForUpcomingCalendarWeekDates(numberOfEntries)

	go keepWeekNumberIconUpToDate()
	go quitOnMenu()
}

func addMenuItemsForUpcomingCalendarWeekDates(numberOfEntries int) {
	for i := 0; i < numberOfEntries; i++ {
		index := i
		menus = append(menus, systray.AddMenuItem("refresh", ""))
		go func() {
			for {
				<-menus[index].ClickedCh
				_, dateToGo := offsetCalendarWeekToDate(index)
				goToDate(dateToGo)
			}
		}()
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
	quitMenuItem := systray.AddMenuItem(fmt.Sprintf("Beenden (%s - %s)", semVer, sha1ver), "Beendet die Applikation")
	<-quitMenuItem.ClickedCh
	systray.Quit()
}

func updateIconAndTooltip(weekNumber int) {
	systray.SetIcon(generateImage(weekNumber))
	systray.SetTooltip(fmt.Sprintf("Aktuelle Kalenderwoche: %d", weekNumber))

	refreshUpcomingCalendarWeekItems()
}

func refreshUpcomingCalendarWeekItems() {
	for index, _ := range menus {
		week, startDate := offsetCalendarWeekToDate(index)

		text := fmt.Sprintf("KW %d: %s", week, monday.Format(startDate, "02. January 2006", monday.LocaleDeDE))
		menus[index].SetTitle(text)
	}
}
