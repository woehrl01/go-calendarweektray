// +build windows

package main

import (
	"golang.org/x/sys/windows"
)

var (
	u32                 = windows.NewLazySystemDLL("User32.dll")
	pSetProcessDPIAware = u32.NewProc("SetProcessDPIAware")
)

func enableDpiAwareness() {
	pSetProcessDPIAware.Call()
}
