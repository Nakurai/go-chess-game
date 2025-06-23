package main

func (b *Board) LoadRegularStart() {
	// place pieces
	b.board[f_a][r_1] = w_rook
	b.board[f_b][r_1] = w_knight
	b.board[f_c][r_1] = w_bishop
	b.board[f_d][r_1] = w_queen
	b.board[f_e][r_1] = w_king
	b.king_position[white] = []int{int(f_e), int(r_1)}
	b.board[f_f][r_1] = w_bishop
	b.board[f_g][r_1] = w_knight
	b.board[f_h][r_1] = w_rook

	b.board[f_a][r_8] = b_rook
	b.board[f_b][r_8] = b_knight
	b.board[f_c][r_8] = b_bishop
	b.board[f_d][r_8] = b_queen
	b.board[f_e][r_8] = b_king
	b.king_position[black] = []int{int(f_e), int(r_8)}
	b.board[f_f][r_8] = b_bishop
	b.board[f_g][r_8] = b_knight
	b.board[f_h][r_8] = b_rook

	for file_index := f_a; file_index <= f_h; file_index++ {
		b.board[file_index][r_2] = w_pawn
		b.board[file_index][r_7] = b_pawn
	}
}
