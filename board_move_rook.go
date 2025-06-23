package main

func (b *Board) GetRookMove(file_index Col, rank_index Rank) []Move {

	return b.GetIncrementalMove(file_index, rank_index, [][]int{
		[]int{0, 1},
		[]int{0, -1},
		[]int{1, 0},
		[]int{-1, 0},
	})
}
