package main

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func loadTileImages() map[string]*ebiten.Image {
	imgs := make(map[string]*ebiten.Image)
	for name, path := range assets_path {
		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			log.Fatalf("failed to load %s: %v", path, err)
		}
		imgs[name] = img
	}
	return imgs
}

var img_map map[string]*ebiten.Image
var tile_size int = 64

func init() {
	img_map = loadTileImages()
}

func main() {

	game := InitGame()

	ebiten.SetWindowSize(1000, 800)
	// ebiten.SetFullscreen(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("X Chess")
	err := ebiten.RunGame(game)
	if err != nil && err.Error() != "exit" {
		log.Fatal(err)
	} else {
		fmt.Println("done")
	}

}
