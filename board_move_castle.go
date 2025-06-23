package main

func (b *Board) GetCastleMoves() []Move {
	possible_moves := []Move{}

	can_long_castle := b.can_w_long_castle
	can_short_castle := b.can_w_short_castle
	if b.turn == black {
		can_long_castle = b.can_b_long_castle
		can_short_castle = b.can_b_short_castle
	}

	if !can_long_castle && !can_short_castle {
		return possible_moves
	}

	file_index := f_e
	rank_index := r_1
	if b.turn == black {
		rank_index = r_8
	}

	piece := w_king
	if b.turn == black {
		piece = b_king
	}

	if b.board[file_index][rank_index] != piece {
		return possible_moves
	}

	opponent_board := b.DeepCopy()
	opponent_board.SwitchTurn()
	opponent_board.ComputeAllPossibleMoves(false, false)

	if can_long_castle {
		castle_is_empty := b.board[f_d][rank_index] == empty && b.board[f_c][rank_index] == empty

		if castle_is_empty {
			files_to_check := []Col{f_e, f_d, f_c, f_b}
			castle_is_under_attack := false
			for _, move := range opponent_board.possible_moves {
				for _, file_to_check := range files_to_check {
					if move.end_file == file_to_check && move.end_rank == rank_index {
						castle_is_under_attack = true
						break
					}
				}
				if castle_is_under_attack {
					break
				}
			}

			if !castle_is_under_attack {
				new_file := f_c
				possible_moves = append(possible_moves, Move{
					start_file:          file_index,
					start_rank:          rank_index,
					end_file:            new_file,
					end_rank:            rank_index,
					piece:               piece,
					turn_to_queen:       false,
					is_capture:          false,
					is_pawn_double_move: false,
					is_long_castle:      true,
					is_short_castle:     false,
				})
			}

		}

	}

	if can_short_castle {
		castle_is_empty := b.board[f_f][rank_index] == empty && b.board[f_g][rank_index] == empty

		if castle_is_empty {
			castle_is_under_attack := false
			files_to_check := []Col{f_e, f_f, f_g}
			for _, move := range opponent_board.possible_moves {
				for _, file_to_check := range files_to_check {
					if move.end_file == file_to_check && move.end_rank == rank_index {
						castle_is_under_attack = true
						break
					}
				}
				if castle_is_under_attack {
					break
				}
			}

			if !castle_is_under_attack {
				new_file := f_g
				possible_moves = append(possible_moves, Move{
					start_file:          file_index,
					start_rank:          rank_index,
					end_file:            new_file,
					end_rank:            rank_index,
					piece:               piece,
					turn_to_queen:       false,
					is_capture:          false,
					is_pawn_double_move: false,
					is_long_castle:      false,
					is_short_castle:     true,
				})
			}

		}

	}

	return possible_moves
}
