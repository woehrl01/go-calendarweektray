package main

import (
	"fmt"
	"time"

	"github.com/getlantern/systray"
	"github.com/goodsign/monday"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
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

func GoToDate(dateToGo time.Time) {
	date := dateToGo.Format("02/01/2006")
	ole.CoInitialize(0)
	defer ole.CoUninitialize()
	unknown, _ := oleutil.CreateObject("Outlook.Application")
	defer unknown.Release()
	outlook, _ := unknown.QueryInterface(ole.IID_IDispatch)
	defer outlook.Release()

	ns := oleutil.MustCallMethod(outlook, "GetNamespace", "MAPI").ToIDispatch()
	calendar := oleutil.MustCallMethod(ns, "GetDefaultFolder", 9).ToIDispatch()
	explorer := oleutil.MustCallMethod(outlook, "ActiveExplorer").ToIDispatch()
	if explorer == nil {
		oleutil.MustCallMethod(calendar, "Display")
		explorer = oleutil.MustCallMethod(outlook, "ActiveExplorer").ToIDispatch()
	}

	oleutil.MustCallMethod(explorer, "CurrentFolder", calendar)
	view := oleutil.MustGetProperty(explorer, "CurrentView").ToIDispatch()
	oleutil.MustCallMethod(view, "GoToDate", date).ToIDispatch()
	oleutil.MustCallMethod(explorer, "Activate")
}

func addMenuItemsForUpcomingCalendarWeekDates() {
	for i := 0; i < 15; i++ {
		var index = i
		menus = append(menus, systray.AddMenuItem("refresh", ""))
		go func() {
			for {
				<-menus[index].ClickedCh
				_, dateToGo := offsetCalendarWeekToDate(index)
				GoToDate(dateToGo)
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
	for index, _ := range menus {
		week, startDate := offsetCalendarWeekToDate(index)

		text := fmt.Sprintf("KW %d: %s", week, monday.Format(startDate, "02. January 2006", monday.LocaleDeDE))
		menus[index].SetTitle(text)
	}
}
