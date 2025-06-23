package main

type Board struct {
	board              [8][8]Piece
	turn               Player
	opponent           Player
	move_50            int
	can_w_long_castle  bool
	can_w_short_castle bool
	can_b_long_castle  bool
	can_b_short_castle bool
	en_passant_file    Col
	en_passant_rank    Rank
	possible_moves     []Move
	king_position      map[Player][]int
}

func (b *Board) Init() {

	b.turn = white
	b.opponent = black
	b.move_50 = 0
	b.can_w_long_castle = true
	b.can_w_short_castle = true
	b.can_b_long_castle = true
	b.can_b_short_castle = true
	b.en_passant_file = f_none
	b.en_passant_rank = r_none
	b.possible_moves = []Move{}
	b.king_position = map[Player][]int{
		white: []int{int(f_none), int(r_none)},
		black: []int{int(f_none), int(r_none)},
	}

	// initialize all pieces to empty
	for file_index := f_a; file_index < f_h+1; file_index++ {
		for rank_index := r_1; rank_index < r_8+1; rank_index++ {
			b.board[file_index][rank_index] = empty
		}
	}

}

func (b *Board) SwitchTurn() {
	if b.turn == white {
		b.turn = black
		b.opponent = white
	} else {
		b.turn = white
		b.opponent = black
	}
}

func (b *Board) ComputeAllPossibleMoves(do_castle bool, check_king bool) {
	// fmt.Println("==========================")
	// fmt.Println("Computing moves for board:")
	// fmt.Println(b)

	all_possible_moves := []Move{}
	for file_index := f_a; file_index < f_h+1; file_index++ {
		for rank_index := r_1; rank_index < r_8+1; rank_index++ {
			piece := b.board[file_index][rank_index]
			if piece == empty {
				continue
			}
			if piece_to_player[piece] != b.turn {
				continue
			}
			piece_type := piece_to_piece_type[piece]
			switch piece_type {
			case pawn:
				possible_move := b.GetPawnMove(file_index, rank_index)
				all_possible_moves = append(all_possible_moves, possible_move...)
				possible_move = b.GetPawnCapture(file_index, rank_index)
				all_possible_moves = append(all_possible_moves, possible_move...)
			case knight:
				possible_move := b.GetKnightMove(file_index, rank_index)
				all_possible_moves = append(all_possible_moves, possible_move...)
			case rook:
				possible_move := b.GetRookMove(file_index, rank_index)
				all_possible_moves = append(all_possible_moves, possible_move...)
			case bishop:
				possible_move := b.GetBishopMove(file_index, rank_index)
				all_possible_moves = append(all_possible_moves, possible_move...)
			case queen:
				possible_move := b.GetRookMove(file_index, rank_index)
				all_possible_moves = append(all_possible_moves, possible_move...)
				possible_move = b.GetBishopMove(file_index, rank_index)
				all_possible_moves = append(all_possible_moves, possible_move...)
			case king:
				possible_move := b.GetKingMove(file_index, rank_index)
				all_possible_moves = append(all_possible_moves, possible_move...)
			}
		}
	}

	// adding en passant moves
	en_passant_possible_move := b.GetPawnEnPassant()
	all_possible_moves = append(all_possible_moves, en_passant_possible_move...)

	// adding castling moves
	if do_castle {
		en_passant_possible_move := b.GetCastleMoves()
		all_possible_moves = append(all_possible_moves, en_passant_possible_move...)
	}

	valid_moves := []Move{}
	if check_king {
		// checking if any of those moves puts the king in check
		for _, move := range all_possible_moves {
			board_copy := b.DeepCopy()
			board_copy.DoMove(move)
			board_copy.ComputeAllPossibleMoves(true, false)

			// fmt.Printf("after move %s, possible moves: %v\n", move, board_copy.possible_moves)

			king_in_check := false
			// the move we made above may have been a king move so we need to recalculate the king's position

			king_file := Col(board_copy.king_position[board_copy.opponent][0])
			king_rank := Rank(board_copy.king_position[board_copy.opponent][1])
			// fmt.Printf("king position is %s%s\n", file_to_string[king_file], rank_to_string[king_rank])
			for _, opponent_move := range board_copy.possible_moves {
				if opponent_move.end_file == king_file && opponent_move.end_rank == king_rank {
					king_in_check = true
					break
				}
			}
			if !king_in_check {
				valid_moves = append(valid_moves, move)
			}
		}
	} else {
		valid_moves = append([]Move{}, all_possible_moves...)
	}

	// fmt.Println("valid moves are:")
	// fmt.Println(valid_moves)
	// fmt.Println("===============================")
	b.possible_moves = valid_moves
}

func (b *Board) DoMove(move Move) {

	b.board[move.end_file][move.end_rank] = b.board[move.start_file][move.start_rank]
	// checking if the move is a pawn reaching the opposite rank
	if move.turn_to_queen {
		new_queen := w_queen
		if b.turn == black {
			new_queen = b_queen
		}
		b.board[move.end_file][move.end_rank] = new_queen
	}
	b.board[move.start_file][move.start_rank] = empty

	// if the move is en passant, we need to capture the passed piece
	if move.is_en_passant {
		b.board[b.en_passant_file][b.en_passant_rank] = empty
	}

	// checking if the move is a pawn double jump
	if move.is_pawn_double_move {
		b.en_passant_file = move.end_file
		b.en_passant_rank = move.end_rank
	} else {
		b.en_passant_file = f_none
		b.en_passant_rank = r_none
	}

	// updating king position
	if piece_to_piece_type[move.piece] == king {
		b.king_position[b.turn][0] = int(move.end_file)
		b.king_position[b.turn][1] = int(move.end_rank)

		// disabling castling rights
		if b.turn == white {
			b.can_w_long_castle = false
			b.can_w_short_castle = false
		} else {
			b.can_b_long_castle = false
			b.can_b_short_castle = false
		}

	}

	// if the move is a rook move, disabling castling rights
	if piece_to_piece_type[move.piece] == rook {
		starting_rook_rank := r_1
		if b.turn == black {
			starting_rook_rank = r_8
		}

		// checking long castle first
		starting_rook_file := f_a
		if move.start_file == starting_rook_file && move.start_rank == starting_rook_rank {
			if b.turn == white {
				b.can_w_long_castle = false
			} else {
				b.can_b_long_castle = false
			}
		}

		// then short castle
		starting_rook_file = f_h
		if move.start_file == starting_rook_file && move.start_rank == starting_rook_rank {
			if b.turn == white {
				b.can_w_short_castle = false
			} else {
				b.can_b_short_castle = false
			}
		}

	}

	// checking if the move is a castle move and updating rook position if so
	if move.is_short_castle {
		rook_rank := r_1
		if b.turn == black {
			rook_rank = r_8
		}
		b.board[f_f][rook_rank] = b.board[f_h][rook_rank]
		b.board[f_h][rook_rank] = empty
	}
	if move.is_long_castle {
		rook_rank := r_1
		if b.turn == black {
			rook_rank = r_8
		}
		b.board[f_d][rook_rank] = b.board[f_a][rook_rank]
		b.board[f_a][rook_rank] = empty
	}

	// @TODO check 50 move rule and implement here

	// starting opponent's turn
	b.SwitchTurn()

}

/*
This function take a direction array and will check all the squares from the original square to those directions
*/
func (b *Board) GetIncrementalMove(file_index Col, rank_index Rank, directions [][]int) []Move {
	possible_moves := []Move{}
	piece := b.board[file_index][rank_index]

	for _, direction := range directions {
		for square_cpt := 1; square_cpt <= 8; square_cpt++ {
			new_piece, new_file, new_rank := b.GetNewSquarePiece(file_index, rank_index, direction[0]*square_cpt, direction[1]*square_cpt)
			// dest is off board
			if new_piece == off_board {
				break
			}
			// dest is occupied by another piece of the current player
			if new_piece != empty && piece_to_player[new_piece] == b.turn {
				break
			}

			is_capture := false
			if new_piece != empty {
				is_capture = true
			}
			possible_moves = append(possible_moves, Move{
				start_file:          file_index,
				start_rank:          rank_index,
				end_file:            new_file,
				end_rank:            new_rank,
				piece:               piece,
				turn_to_queen:       false,
				is_capture:          is_capture,
				is_pawn_double_move: false,
				is_long_castle:      false,
				is_short_castle:     false,
			})
			// if the piece captured another piece, it cannot go beyond
			if is_capture {
				break
			}

		}

	}

	return possible_moves
}

func (b *Board) GetNewSquarePiece(file Col, rank Rank, file_dir int, rank_dir int) (Piece, Col, Rank) {
	new_file := Col(int(file) + file_dir)
	new_rank := Rank(int(rank) + rank_dir)
	new_file_is_on_board := new_file >= f_a && new_file <= f_h
	new_rank_is_on_board := new_rank >= r_1 && new_rank <= r_8
	if !new_file_is_on_board || !new_rank_is_on_board {
		return off_board, new_file, new_rank
	}
	return b.board[new_file][new_rank], new_file, new_rank
}

/*
This function take a direction array and will check all the squares at those ending positions, but not the intermediary squares
*/
func (b *Board) GetDiscreteMove(file_index Col, rank_index Rank, directions [][]int) []Move {
	possible_moves := []Move{}
	piece := b.board[file_index][rank_index]

	for _, direction := range directions {
		new_piece, new_file, new_rank := b.GetNewSquarePiece(file_index, rank_index, direction[0], direction[1])

		// dest is off board
		if new_piece == off_board {
			continue
		}
		// dest is occupied by another piece of the current player
		if new_piece != empty && piece_to_player[new_piece] == b.turn {
			continue
		}

		is_capture := false
		if new_piece != empty {
			is_capture = true
		}
		possible_moves = append(possible_moves, Move{
			start_file:          file_index,
			start_rank:          rank_index,
			end_file:            new_file,
			end_rank:            new_rank,
			piece:               piece,
			turn_to_queen:       false,
			is_capture:          is_capture,
			is_pawn_double_move: false,
			is_long_castle:      false,
			is_short_castle:     false,
		})

	}

	return possible_moves
}

func (b *Board) DeepCopy() *Board {
	new_board := Board{}
	new_board.Init()

	new_board.board = DeepCopy8x8Piece(b.board)
	new_board.turn = b.turn
	new_board.opponent = b.opponent
	new_board.move_50 = b.move_50
	new_board.can_w_long_castle = b.can_w_long_castle
	new_board.can_w_short_castle = b.can_w_short_castle
	new_board.can_b_long_castle = b.can_b_long_castle
	new_board.can_b_short_castle = b.can_b_short_castle
	new_board.en_passant_file = b.en_passant_file
	new_board.en_passant_rank = b.en_passant_rank
	new_board.possible_moves = []Move{}
	new_board.king_position[white][0] = b.king_position[white][0]
	new_board.king_position[white][1] = b.king_position[white][1]
	new_board.king_position[black][0] = b.king_position[black][0]
	new_board.king_position[black][1] = b.king_position[black][1]

	return &new_board

}

func (b Board) String() string {
	res := ""
	for rank_index := r_8; rank_index >= r_1; rank_index-- {
		res += "\n"
		for file_index := f_a; file_index <= f_h; file_index++ {
			piece := b.board[file_index][rank_index]
			res += " " + piece_to_string[piece] + " "
		}
	}
	res += "\n\nPlayer Turn: " + player_to_string[b.turn] + "\n"
	return res
}
