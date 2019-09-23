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

	mQuit := systray.AddMenuItem("Beenden", "Beendet die Applikation")

	ticker := time.NewTicker(10 * time.Minute)

	go func() {
		updateWeekNumber()
		for {
			select {
			case <-ticker.C:
				updateWeekNumber()
				break
			case <-mQuit.ClickedCh:
				ticker.Stop()
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {

}

func updateWeekNumber() {
	_, week := time.Now().UTC().ISOWeek()

	systray.SetIcon(generateImage(week))
	systray.SetTooltip(fmt.Sprintf("Aktuelle Kalenderwoche: %d", week))
}

func generateImage(week int) []byte {

	fontPath, err := findfont.Find("segoeui.ttf")
	if err != nil {
		panic(err)
	}

	const iconSize = 64
	const fontSize = 50

	dc := gg.NewContext(iconSize, iconSize)
	dc.SetRGB(1, 1, 1)
	if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored(fmt.Sprintf("%d", week), iconSize/2, iconSize/2, 0.5, 0.5)
	img := dc.Image()

	var b bytes.Buffer
	foo := bufio.NewWriter(&b)
	ico.Encode(foo, img)
	foo.Flush()
	return b.Bytes()
}
