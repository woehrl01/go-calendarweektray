package main

import (
	"bufio"
	"bytes"
	"fmt"
	"time"

	ico "github.com/biessek/golang-ico"
	findfont "github.com/flopp/go-findfont"
	"github.com/fogleman/gg"
	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Kalenderwoche")

	go keepWeekNumberIconUpToDate()
	go quitOnMenu()
}

func onExit() {

}

func keepWeekNumberIconUpToDate(){
	for ticker := time.NewTicker(10 * time.Minute) ; ; <-ticker.C {
		updateWeekNumber()
	}
}

func quitOnMenu(){
	mQuit := systray.AddMenuItem("Beenden", "Beendet die Applikation")
	<-mQuit.ClickedCh
	systray.Quit()
}

func updateWeekNumber() {
	_, week := time.Now().UTC().ISOWeek()

	systray.SetIcon(generateImage(week))
	systray.SetTooltip(fmt.Sprintf("Aktuelle Kalenderwoche: %d", week))
}

func generateImage(week int) []byte {

	const iconSize = 64
	const fontSize = 50

	fontPath, err := findfont.Find("segoeui.ttf")
	if err != nil {
		panic(err)
	}

	dc := gg.NewContext(iconSize, iconSize)
	if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
		panic(err)
	}

	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(fmt.Sprintf("%d", week), iconSize/2, iconSize/2, 0.5, 0.5)
	img := dc.Image()

	var b bytes.Buffer
	foo := bufio.NewWriter(&b)
	ico.Encode(foo, img)
	foo.Flush()
	return b.Bytes()
}
