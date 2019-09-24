package main

import (
	"bufio"
	"bytes"
	"fmt"

	ico "github.com/biessek/golang-ico"
	findfont "github.com/flopp/go-findfont"
	"github.com/fogleman/gg"
)

func generateImage(weekNumber int) []byte {
	const iconSize = 64
	const fontSize = 50
	
	dc := gg.NewContext(iconSize, iconSize)

	setFont(dc, "segoeui.ttf", fontSize);

	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(fmt.Sprintf("%d", weekNumber), iconSize/2, iconSize/2, 0.5, 0.5)

	return writeContextToByteArray(dc)
}

func setFont(dc *gg.Context, fontName string, fontSize float64){
	fontPath, err := findfont.Find(fontName)
	if err != nil {
		panic(err)
	}
	
	if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
		panic(err)
	}
}

func writeContextToByteArray(dc *gg.Context)  []byte{
	var b bytes.Buffer
	foo := bufio.NewWriter(&b)
	ico.Encode(foo, dc.Image())
	foo.Flush()
	return b.Bytes()
}