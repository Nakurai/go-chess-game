package main

import "github.com/hajimehoshi/ebiten/v2"

var piece_scale float64 = 40.0 / 18.0

func (g *Game) GameBoardDraw(screen *ebiten.Image) {
	for rank_index := r_8; rank_index >= r_1; rank_index-- {
		for file_index := f_a; file_index <= f_h; file_index++ {
			square := g.visible_board.Squares[file_index][rank_index]
			img_to_draw := img_map[square.sq_type]
			op := ebiten.DrawImageOptions{}
			if square.sq_status != "neutral" {
				op.ColorScale.Scale(1, 0, 0, 0.5)
			}
			op.GeoM.Translate(square.sq_x+100, square.sq_y+60)
			screen.DrawImage(img_to_draw, &op)

			piece_op := ebiten.DrawImageOptions{}
			piece_op.GeoM.Scale(piece_scale, piece_scale)
			piece_op.GeoM.Translate(square.sq_x+100+12, square.sq_y+60+12)
			piece := g.board.board[file_index][rank_index]
			if piece != empty {
				piece_name := piece_to_string[piece]
				img_to_draw = img_map[piece_name]
				screen.DrawImage(img_to_draw, &piece_op)
			}
		}
	}
}
