// +build windows

package main

import (
	"time"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func goToDate(dateToGo time.Time) {
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
		/*Start outllok if not started*/
		oleutil.MustCallMethod(calendar, "Display")
		explorer = oleutil.MustCallMethod(outlook, "ActiveExplorer").ToIDispatch()
	}

	oleutil.MustCallMethod(explorer, "CurrentFolder", calendar)
	view := oleutil.MustGetProperty(explorer, "CurrentView").ToIDispatch()
	oleutil.MustCallMethod(view, "GoToDate", date).ToIDispatch()
	oleutil.MustCallMethod(explorer, "Activate")
}
