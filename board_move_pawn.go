package main

func (b *Board) GetPawnMove(file_index Col, rank_index Rank) []Move {
	possible_moves := []Move{}
	piece := b.board[file_index][rank_index]

	// regular forward move
	direction := 1
	if b.turn == black {
		direction = -1
	}
	new_piece, new_file, new_rank := b.GetNewSquarePiece(file_index, rank_index, 0, direction)
	if new_piece == empty {
		turn_to_queen := new_rank == r_1 || new_rank == r_8
		possible_moves = append(possible_moves, Move{
			start_file:          file_index,
			start_rank:          rank_index,
			end_file:            new_file,
			end_rank:            new_rank,
			piece:               piece,
			turn_to_queen:       turn_to_queen,
			is_capture:          false,
			is_pawn_double_move: false,
			is_long_castle:      false,
			is_short_castle:     false,
		})

		// if the square in front of the pawn is empty, then
		// it may be possible to jump two squares if the pawn
		// is on the initial rank
		first_pawn_move := (b.turn == white && rank_index == r_2) || (b.turn == black && rank_index == r_7)
		if first_pawn_move {
			new_piece, new_file, new_rank = b.GetNewSquarePiece(file_index, rank_index, 0, 2*direction)
			if new_piece == empty {
				turn_to_queen := new_rank == r_1 || new_rank == r_8
				possible_moves = append(possible_moves, Move{
					start_file:          file_index,
					start_rank:          rank_index,
					end_file:            new_file,
					end_rank:            new_rank,
					piece:               piece,
					turn_to_queen:       turn_to_queen,
					is_capture:          false,
					is_pawn_double_move: true,
					is_long_castle:      false,
					is_short_castle:     false,
				})
			}
		}
	}

	return possible_moves
}

func (b *Board) GetPawnEnPassant() []Move {
	possible_moves := []Move{}

	if b.en_passant_file != f_none && b.en_passant_rank != r_none {
		direction := 1
		if b.turn == black {
			direction = -1
		}
		end_rank := Rank(int(b.en_passant_rank) + direction)
		new_piece, _, _ := b.GetNewSquarePiece(b.en_passant_file, b.en_passant_rank, -1, 0)
		if piece_to_piece_type[new_piece] == pawn && piece_to_player[new_piece] == b.turn {
			possible_moves = append(possible_moves, Move{
				start_file:          b.en_passant_file - 1,
				start_rank:          b.en_passant_rank,
				end_file:            b.en_passant_file,
				end_rank:            end_rank,
				piece:               new_piece,
				turn_to_queen:       false,
				is_capture:          true,
				is_pawn_double_move: false,
				is_long_castle:      false,
				is_short_castle:     false,
				is_en_passant:       true,
			})
		}
		new_piece, _, _ = b.GetNewSquarePiece(b.en_passant_file, b.en_passant_rank, 1, 0)
		if piece_to_piece_type[new_piece] == pawn && piece_to_player[new_piece] == b.turn {
			possible_moves = append(possible_moves, Move{
				start_file:          b.en_passant_file + 1,
				start_rank:          b.en_passant_rank,
				end_file:            b.en_passant_file,
				end_rank:            end_rank,
				piece:               new_piece,
				turn_to_queen:       false,
				is_capture:          true,
				is_pawn_double_move: false,
				is_long_castle:      false,
				is_short_castle:     false,
				is_en_passant:       true,
			})
		}

	}
	return possible_moves
}

func (b *Board) GetPawnCapture(file_index Col, rank_index Rank) []Move {
	possible_moves := []Move{}
	piece := b.board[file_index][rank_index]

	// regular forward move
	direction := 1
	if b.turn == black {
		direction = -1
	}

	// first diagonal
	new_piece, new_file, new_rank := b.GetNewSquarePiece(file_index, rank_index, 1, direction)
	if new_piece != empty && piece_to_player[new_piece] == b.opponent {
		turn_to_queen := new_rank == r_1 || new_rank == r_8
		possible_moves = append(possible_moves, Move{
			start_file:          file_index,
			start_rank:          rank_index,
			end_file:            new_file,
			end_rank:            new_rank,
			piece:               piece,
			turn_to_queen:       turn_to_queen,
			is_capture:          true,
			is_pawn_double_move: false,
			is_long_castle:      false,
			is_short_castle:     false,
		})
	}

	// second diagonal
	new_piece, new_file, new_rank = b.GetNewSquarePiece(file_index, rank_index, -1, direction)
	if new_piece != empty && piece_to_player[new_piece] == b.opponent {
		turn_to_queen := new_rank == r_1 || new_rank == r_8
		possible_moves = append(possible_moves, Move{
			start_file:          file_index,
			start_rank:          rank_index,
			end_file:            new_file,
			end_rank:            new_rank,
			piece:               piece,
			turn_to_queen:       turn_to_queen,
			is_capture:          true,
			is_pawn_double_move: false,
			is_long_castle:      false,
			is_short_castle:     false,
		})
	}

	return possible_moves
}
