package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	board         Board
	visible_board VisibleBoard
	ui            string
	options       GameOptions
	menu          GameMenu
}

func InitGame() *Game {
	board := Board{}
	board.Init()
	board.LoadRegularStart()
	board.ComputeAllPossibleMoves(true, true)
	visible_board := VisibleBoard{}
	visible_board.Init()
	options := GameOptions{
		Mode: "pvp",
	}
	menu := GameMenu{
		SelectedIndex: 0,
		Items: []GameMenuItem{
			GameMenuItem{
				Label: "Player vs Player",
				X:     100,
				Y:     60,
			},
			GameMenuItem{
				Label: "Player vs AI",
				X:     100,
				Y:     140,
			},
			GameMenuItem{
				Label: "Exit",
				X:     100,
				Y:     220,
			},
		},
	}
	game := &Game{
		board:         board,
		visible_board: visible_board,
		ui:            "menu",
		options:       options,
		menu:          menu,
	}
	return game
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) && g.ui != "menu" {
		g.ui = "menu"
		return nil
	}

	switch g.ui {
	case "board":
		return g.GameBoardUpdate()
	case "menu":
		return g.GameMenuUpdate()
	default:
		fmt.Println("Error!")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.ui {
	case "board":
		g.GameBoardDraw(screen)
	case "menu":
		g.GameMenuDraw(screen)
	default:
		fmt.Println("Error!")

	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}
