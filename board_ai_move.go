package main

import "math/rand/v2"

func (b *Board) GetAiMove() Move {
	nextMoveIndex := rand.IntN(len(b.possible_moves))
	return b.possible_moves[nextMoveIndex]
}
