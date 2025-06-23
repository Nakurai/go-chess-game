package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) GameMenuUpdate() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		if g.menu.SelectedIndex > 0 {
			g.menu.SelectedIndex -= 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		if g.menu.SelectedIndex < float64(len(g.menu.Items)-1) {
			g.menu.SelectedIndex += 1
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) && g.ui == "menu" {
		return fmt.Errorf("exit")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch g.menu.SelectedIndex {
		case 0:
			g.options.Mode = "pvp"
			g.ui = "board"
		case 1:
			g.options.Mode = "pvai"
			g.ui = "board"
		case 2:
			return fmt.Errorf("exit")
		}
	}
	return nil
}
