package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const boardOffsetX = 100
const boardOffsetY = 60

func (g *Game) GameBoardUpdate() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		fmt.Println(g.board)
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		// Adjust for offset
		boardX := x - boardOffsetX
		boardY := y - boardOffsetY

		// Check if click is inside the board area
		selection_on_board := true
		if boardX < 0 || boardY < 0 || boardX >= tile_size*8 || boardY >= tile_size*8 {
			selection_on_board = false
		}

		tile_file := Col(boardX / tile_size)
		tile_rank := Rank(7 - (boardY / tile_size)) // still flipped if your board is "white on bottom"

		piece := g.board.board[tile_file][tile_rank]

		for _, move := range g.board.possible_moves {
			if move.Is(g.visible_board.SelectedFile, g.visible_board.SelectedRank, tile_file, tile_rank) {
				g.board.DoMove(move)
				g.board.ComputeAllPossibleMoves(true, true)
				g.visible_board.ResetAllTiles()

				if len(g.board.possible_moves) == 0 {
					// game is over!

				} else {
					// if the user play against the ai, we need to pick the next move
					if g.options.Mode == "pvai" {
						nextMove := g.board.GetAiMove()
						g.board.DoMove(nextMove)
						g.board.ComputeAllPossibleMoves(true, true)
					}
				}

				return nil
			}
		}

		if selection_on_board {
			tile_already_selected := g.visible_board.SelectedFile == tile_file && g.visible_board.SelectedRank == tile_rank
			selection_is_empty := piece == empty
			if !tile_already_selected && !selection_is_empty {
				selection_is_right_turn := g.board.turn == piece_to_player[piece]
				if selection_is_right_turn {
					g.visible_board.SelectTile(tile_file, tile_rank)
					for _, move := range g.board.possible_moves {
						if move.start_file == tile_file && move.start_rank == tile_rank {
							g.visible_board.TargetTile(move.end_file, move.end_rank)
						}
					}
				}
			} else {
				g.visible_board.ResetAllTiles()
			}
		}
	}

	return nil
}
