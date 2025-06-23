package main

func (b *Board) GetKnightMove(file_index Col, rank_index Rank) []Move {
	return b.GetDiscreteMove(file_index, rank_index, [][]int{
		[]int{1, 2},
		[]int{-1, 2},
		[]int{2, 1},
		[]int{2, -1},
		[]int{1, -2},
		[]int{-1, -2},
		[]int{-2, 1},
		[]int{-2, -1},
	})
}
