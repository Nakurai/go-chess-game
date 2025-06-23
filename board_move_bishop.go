package main

func (b *Board) GetBishopMove(file_index Col, rank_index Rank) []Move {

	return b.GetIncrementalMove(file_index, rank_index, [][]int{
		[]int{1, 1},
		[]int{1, -1},
		[]int{-1, -1},
		[]int{-1, 1},
	})
}
