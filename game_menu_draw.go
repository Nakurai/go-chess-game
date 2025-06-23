package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	mplusFaceSource *text.GoTextFaceSource
	normalFont      *text.GoTextFace
	gray            = color.RGBA{0x80, 0x80, 0x80, 0xff}
)

const (
	normalFontSize = 24
	bigFontSize    = 48
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
	normalFont = &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   normalFontSize,
	}
}

func (g *Game) GameMenuDraw(screen *ebiten.Image) {

	for itemIndex, menuItem := range g.menu.Items {

		if g.menu.SelectedIndex == float64(itemIndex) {
			w, h := text.Measure(menuItem.Label, normalFont, 0)
			vector.DrawFilledRect(screen, float32(menuItem.X)-5, float32(menuItem.Y)-5, float32(w)+10, float32(h)+10, gray, false)
		}

		op := &text.DrawOptions{}
		op.GeoM.Translate(menuItem.X, menuItem.Y)
		op.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, menuItem.Label, normalFont, op)
	}
}
