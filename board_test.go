package main

import (
	"testing"
)

func TestPieceLocation(t *testing.T) {
	b := Board{}
	b.Init()
	b.LoadRegularStart()

	if b.board[f_a][r_1] != w_rook {
		t.Errorf("a1 should be R (%d), got %d", w_rook, b.board[f_a][r_1])
	}
	if b.board[f_b][r_1] != w_knight {
		t.Errorf("b1 should be N (%d), got %d", w_knight, b.board[f_b][r_1])
	}
}

func TestWhitePawnMove(t *testing.T) {
	b := Board{}
	b.Init()

	/*
		a  b  c  d  e  f  g  h
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  r  .  .  .
		.  P  .  .  .  .  .  r
		P  .  .  .  P  .  .  P
		.  .  .  .  .  .  .  .
	*/
	b.board[f_a][r_2] = w_pawn // can move 1 and 2 squares
	b.board[f_b][r_3] = w_pawn // can move 1 square only
	b.board[f_e][r_2] = w_pawn // can only move 1 square (2nd is blocked)
	b.board[f_e][r_4] = b_rook
	b.board[f_h][r_2] = w_pawn // cannot move at all
	b.board[f_h][r_3] = b_rook

	// fmt.Println(b)
	b.ComputeAllPossibleMoves(true, true)

	AssessMoves(t, b, []string{
		"a2a3",
		"a2a4",
		"b3b4",
		"e2e3",
	})

}

func TestWhitePawnCapture(t *testing.T) {
	b := Board{}
	b.Init()

	/*
		a  b  c  d  e  f  g  h
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  p  .  .  .
		r  r  r  r  P  .  r  r
		P  .  P  P  .  .  .  P
		.  .  .  .  .  .  .  .
	*/
	b.board[f_a][r_2] = w_pawn
	b.board[f_a][r_3] = b_rook
	b.board[f_b][r_3] = b_rook
	b.board[f_c][r_2] = w_pawn
	b.board[f_c][r_3] = b_rook
	b.board[f_d][r_2] = w_pawn
	b.board[f_d][r_3] = b_rook
	b.board[f_e][r_3] = w_pawn
	b.board[f_e][r_4] = b_pawn
	b.board[f_f][r_2] = b_rook
	b.board[f_g][r_3] = b_rook
	b.board[f_h][r_2] = w_pawn
	b.board[f_h][r_3] = b_rook

	// fmt.Println(b)
	b.ComputeAllPossibleMoves(true, true)

	AssessMoves(t, b, []string{
		"a2xb3",
		"c2xb3",
		"c2xd3",
		"d2xc3",
		"h2xg3",
	})

}

func TestWhitePawnEnPassant(t *testing.T) {
	b := Board{}
	b.Init()

	/*
		a  b  c  d  e  f  g  h
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  r  .  r  .
		.  .  .  .  P  p  P  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
	*/
	b.en_passant_file = f_f
	b.en_passant_rank = r_5
	b.board[f_e][r_6] = b_rook
	b.board[f_e][r_5] = w_pawn
	b.board[f_f][r_5] = b_pawn
	b.board[f_g][r_6] = b_rook
	b.board[f_g][r_5] = w_pawn

	// fmt.Println(b)
	b.ComputeAllPossibleMoves(true, true)

	AssessMoves(t, b, []string{
		"e5xf6",
		"g5xf6",
	})

}

func TestBlackPawnMove(t *testing.T) {
	b := Board{}
	b.Init()

	/*
		a  b  c  d  e  f  g  h
		.  .  .  .  .  .  .  .
		p  .  .  .  .  p  .  p
		.  p  .  .  .  .  .  R
		.  .  .  .  .  R  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
	*/
	b.SwitchTurn()
	b.board[f_a][r_7] = b_pawn // can move 1 and 2 squares
	b.board[f_b][r_6] = b_pawn // can move 1 square only
	b.board[f_f][r_7] = b_pawn // can only move 1 square (2nd is blocked)
	b.board[f_f][r_5] = w_rook
	b.board[f_h][r_7] = b_pawn // cannot move at all
	b.board[f_h][r_6] = w_rook

	// fmt.Println(b)
	b.ComputeAllPossibleMoves(true, true)

	AssessMoves(t, b, []string{
		"a7a6",
		"a7a5",
		"b6b5",
		"f7f6",
	})

}

func TestBlackPawnCapture(t *testing.T) {
	b := Board{}
	b.Init()
	b.SwitchTurn()

	/*
		a  b  c  d  e  f  g  h
		.  .  .  .  .  .  .  .
		p  .  p  .  p  .  .  p
		R  R  R  R  R  p  R  R
		.  .  .  .  .  P  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
	*/
	b.board[f_a][r_7] = b_pawn
	b.board[f_a][r_6] = w_rook
	b.board[f_b][r_6] = w_rook
	b.board[f_c][r_7] = b_pawn
	b.board[f_c][r_6] = w_rook
	b.board[f_d][r_6] = w_rook
	b.board[f_e][r_7] = b_pawn
	b.board[f_e][r_6] = w_rook
	b.board[f_f][r_6] = b_pawn
	b.board[f_f][r_5] = w_pawn
	b.board[f_g][r_6] = w_rook
	b.board[f_h][r_7] = b_pawn
	b.board[f_h][r_6] = w_rook

	// fmt.Println(b)
	b.ComputeAllPossibleMoves(true, true)

	AssessMoves(t, b, []string{
		"a7xb6",
		"c7xb6",
		"c7xd6",
		"e7xd6",
		"h7xg6",
	})

}

func TestBlackPawnEnPassant(t *testing.T) {
	b := Board{}
	b.Init()

	/*
		a  b  c  d  e  f  g  h
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  p  P  p  .
		.  .  .  .  R  .  R  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
	*/
	b.SwitchTurn()
	b.en_passant_file = f_f
	b.en_passant_rank = r_4
	b.board[f_e][r_3] = w_rook
	b.board[f_e][r_4] = b_pawn
	b.board[f_f][r_4] = w_pawn
	b.board[f_g][r_3] = w_rook
	b.board[f_g][r_4] = b_pawn

	// fmt.Println(b)
	b.ComputeAllPossibleMoves(true, true)

	AssessMoves(t, b, []string{
		"e4xf3",
		"g4xf3",
	})

}

func TestKnightMove(t *testing.T) {
	b := Board{}
	b.Init()

	/*
		a  b  c  d  e  f  g  h
		.  P  .  .  .  .  .  .
		.  P  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		N  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  p  .  .  .  .  .  .
		.  .  .  .    .  .  .
		.  .  .  .  .  .  .  .
	*/
	b.board[f_a][r_5] = w_knight
	b.board[f_b][r_8] = w_pawn
	b.board[f_b][r_7] = w_pawn
	b.board[f_b][r_3] = b_pawn

	// fmt.Println(b)
	b.ComputeAllPossibleMoves(true, true)

	AssessMoves(t, b, []string{
		"a5c6",
		"a5c4",
		"a5xb3",
	})

}

func TestRookMove(t *testing.T) {
	b := Board{}
	b.Init()

	/*
		a  b  c  d  e  f  g  h
		.  .  .  P  .  .  .  .
		.  .  .  P  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  R  .  p  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .    .  .  .
		.  .  .  .  .  .  .  .
	*/
	b.board[f_d][r_5] = w_rook
	b.board[f_d][r_7] = w_pawn
	b.board[f_d][r_8] = w_pawn
	b.board[f_f][r_5] = b_pawn

	// fmt.Println(b)
	b.ComputeAllPossibleMoves(true, true)

	AssessMoves(t, b, []string{
		"d5d6",
		"d5c5",
		"d5b5",
		"d5a5",
		"d5d4",
		"d5d3",
		"d5d2",
		"d5d1",
		"d5e5",
		"d5xf5",
	})

}

func TestBishopMove(t *testing.T) {
	b := Board{}
	b.Init()

	/*
		a  b  c  d  e  f  g  h
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  p  .  .  .
		.  .  .  .  P  .  .  .
		.  p  .  .  .  .  .  .
		.  .  B  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
	*/
	b.board[f_c][r_3] = w_bishop
	b.board[f_b][r_4] = b_pawn
	b.board[f_e][r_5] = w_pawn
	b.board[f_e][r_6] = b_pawn

	// fmt.Println(b)
	b.ComputeAllPossibleMoves(true, true)

	AssessMoves(t, b, []string{
		"c3xb4",
		"c3b2",
		"c3a1",
		"c3d2",
		"c3e1",
		"c3d4",
	})

}

func TestKingMove(t *testing.T) {
	b := Board{}
	b.Init()

	/*
		a  b  c  d  e  f  g  h
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .    .  .  .
		.  .  K  .  .  .  .  .
	*/
	b.board[f_c][r_1] = w_king
	b.king_position[white] = []int{int(f_c), int(r_1)}

	// fmt.Println(b)
	b.ComputeAllPossibleMoves(true, true)

	AssessMoves(t, b, []string{
		"c1b1",
		"c1b2",
		"c1c2",
		"c1d2",
		"c1d1",
	})

}

func TestKingMoveCapture(t *testing.T) {
	b := Board{}
	b.Init()

	/*
		a  b  c  d  e  f  g  h
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  p  .  .  .  .
		.  .  p  P    .  .  .
		.  .  K  .  .  .  .  .
	*/
	b.board[f_c][r_1] = w_king
	b.king_position[white] = []int{int(f_c), int(r_1)}
	b.board[f_c][r_2] = b_pawn
	b.board[f_d][r_2] = w_pawn
	b.board[f_d][r_3] = b_pawn

	// fmt.Println(b)
	b.ComputeAllPossibleMoves(true, true)

	AssessMoves(t, b, []string{
		"c1b2",
	})

}

func TestCastleShortMoveOk(t *testing.T) {
	b := Board{}
	b.Init()

	/*
		a  b  c  d  e  f  g  h
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  P  P  P  .  P
		.  .  .  P  P  P  .  P
		.  .  .  P  K  .  .  R
	*/
	b.board[f_d][r_1] = w_pawn
	b.board[f_d][r_2] = w_pawn
	b.board[f_d][r_3] = w_pawn
	b.board[f_e][r_1] = w_king
	b.board[f_e][r_2] = w_pawn
	b.board[f_e][r_3] = w_pawn
	b.board[f_f][r_2] = w_pawn
	b.board[f_f][r_3] = w_pawn
	b.board[f_h][r_1] = w_rook
	b.board[f_h][r_2] = w_pawn
	b.board[f_h][r_3] = w_pawn

	// fmt.Println(b)
	b.ComputeAllPossibleMoves(true, true)

	AssessMoves(t, b, []string{
		"e1f1",
		"e1g1",
		"h1g1",
		"h1f1",
		"d3d4",
		"e3e4",
		"f3f4",
		"h3h4",
	})
}

func TestCastleShortMoveUnderAttack(t *testing.T) {
	b := Board{}
	b.Init()

	/*
		a  b  c  d  e  f  g  h
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  r  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  P  P  P  .  P
		.  .  .  P  P  P  .  P
		.  .  .  P  K  .  .  R
	*/
	b.board[f_d][r_1] = w_pawn
	b.board[f_d][r_2] = w_pawn
	b.board[f_d][r_3] = w_pawn
	b.board[f_e][r_1] = w_king
	b.board[f_e][r_2] = w_pawn
	b.board[f_e][r_3] = w_pawn
	b.board[f_f][r_2] = w_pawn
	b.board[f_f][r_3] = w_pawn
	b.board[f_g][r_6] = b_rook
	b.board[f_h][r_1] = w_rook
	b.board[f_h][r_2] = w_pawn
	b.board[f_h][r_3] = w_pawn

	// fmt.Println(b)
	b.ComputeAllPossibleMoves(true, true)

	AssessMoves(t, b, []string{
		"e1f1",
		"h1g1",
		"h1f1",
		"d3d4",
		"e3e4",
		"f3f4",
		"h3h4",
	})
}

func TestCastleLongMoveOk(t *testing.T) {
	b := Board{}
	b.Init()

	/*
		a  b  c  d  e  f  g  h
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		P  .  .  P  P  P  .  .
		P  .  .  P  P  P  .  .
		R  .  .  .  K  P  .  .
	*/
	b.board[f_a][r_1] = w_rook
	b.board[f_a][r_2] = w_pawn
	b.board[f_a][r_3] = w_pawn
	b.board[f_d][r_2] = w_pawn
	b.board[f_d][r_3] = w_pawn
	b.board[f_e][r_1] = w_king
	b.board[f_e][r_2] = w_pawn
	b.board[f_e][r_3] = w_pawn
	b.board[f_f][r_1] = w_pawn
	b.board[f_f][r_2] = w_pawn
	b.board[f_f][r_3] = w_pawn

	// fmt.Println(b)
	b.ComputeAllPossibleMoves(true, true)

	AssessMoves(t, b, []string{
		"a1b1",
		"a1c1",
		"a1d1",
		"a3a4",
		"e1d1",
		"e1c1",
		"d3d4",
		"e3e4",
		"f3f4",
	})
}

func TestCastleLongMoveUnderAttack(t *testing.T) {
	b := Board{}
	b.Init()

	/*
		a  b  c  d  e  f  g  h
		.  r  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		P  .  .  P  P  P  .  .
		P  .  .  P  P  P  .  .
		R  .  .  .  K  P  .  .
	*/
	b.board[f_a][r_1] = w_rook
	b.board[f_a][r_2] = w_pawn
	b.board[f_a][r_3] = w_pawn
	b.board[f_b][r_8] = b_rook
	b.board[f_d][r_2] = w_pawn
	b.board[f_d][r_3] = w_pawn
	b.board[f_e][r_1] = w_king
	b.board[f_e][r_2] = w_pawn
	b.board[f_e][r_3] = w_pawn
	b.board[f_f][r_1] = w_pawn
	b.board[f_f][r_2] = w_pawn
	b.board[f_f][r_3] = w_pawn

	// fmt.Println(b)
	b.ComputeAllPossibleMoves(true, true)

	AssessMoves(t, b, []string{
		"a1b1",
		"a1c1",
		"a1d1",
		"a3a4",
		"e1d1",
		"d3d4",
		"e3e4",
		"f3f4",
	})
}

func TestMoveKingCheck(t *testing.T) {
	b := Board{}
	b.Init()

	/*
		a  b  c  d  e  f  g  h
		.  r  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		P  .  .  P  P  P  .  .
		P  .  .  P  P  P  .  .
		R  .  K  .  .  P  .  .
	*/
	b.board[f_a][r_1] = w_rook
	b.board[f_a][r_2] = w_pawn
	b.board[f_a][r_3] = w_pawn
	b.board[f_b][r_8] = b_rook
	b.board[f_d][r_2] = w_pawn
	b.board[f_d][r_3] = w_pawn
	b.board[f_c][r_1] = w_king
	b.king_position[white] = []int{int(f_c), int(r_1)}
	b.board[f_e][r_2] = w_pawn
	b.board[f_e][r_3] = w_pawn
	b.board[f_f][r_1] = w_pawn
	b.board[f_f][r_2] = w_pawn
	b.board[f_f][r_3] = w_pawn

	b.ComputeAllPossibleMoves(true, true)

	for _, move := range b.possible_moves {
		is_c1_b1 := move.Is(f_c, r_1, f_b, r_1)
		is_c1_b2 := move.Is(f_c, r_1, f_b, r_2)
		if is_c1_b1 || is_c1_b2 {
			t.Errorf("Invalid move, the king cannot go to b1 or b2")
		}
	}
}

func TestMoveKingCheck2(t *testing.T) {
	b := Board{}
	b.Init()

	/*
		a  b  c  d  e  f  g  h
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		.  b  .  .  .  .  .  .
		.  .  .  .  .  .  .  .
		P  P  P  P  P  P  P  P
		.  .  .  .  K  .  .  .
	*/
	b.can_w_long_castle = false
	b.can_w_short_castle = false
	b.board[f_a][r_2] = w_pawn
	b.board[f_b][r_2] = w_pawn
	b.board[f_c][r_2] = w_pawn
	b.board[f_d][r_2] = w_pawn
	b.board[f_e][r_2] = w_pawn
	b.board[f_f][r_2] = w_pawn
	b.board[f_g][r_2] = w_pawn
	b.board[f_h][r_2] = w_pawn
	b.board[f_e][r_1] = w_king
	b.king_position[white] = []int{int(f_e), int(r_1)}
	b.board[f_b][r_4] = b_bishop

	b.ComputeAllPossibleMoves(true, true)

	for _, move := range b.possible_moves {
		is_d2_d3 := move.Is(f_d, r_2, f_d, r_3)
		is_d2_d4 := move.Is(f_d, r_2, f_d, r_4)
		if is_d2_d3 || is_d2_d4 {
			t.Errorf("Invalid move, the pawn cannot go to d2 or d3 because the king would be in check")
		}
	}
}

func AssessMoves(t *testing.T, b Board, valid_moves []string) {
	found_moves := map[string]bool{}
	for _, move := range b.possible_moves {
		move_str := move.String()
		found_moves[move_str] = true
	}
	if len(b.possible_moves) != len(valid_moves) {
		t.Errorf("Wrong number of moves (expected %d vs found %d)!\nValid moves are: %v\nFound moves are: %s\n", len(valid_moves), len(found_moves), valid_moves, b.possible_moves)
	}
	for _, valid_move := range valid_moves {
		if !found_moves[valid_move] {
			t.Errorf("Missing move %s!\nValid moves are: %v\nFound moves are: %v\n", valid_move, valid_moves, found_moves)
		}
	}
}

func TestCopyBoard(t *testing.T) {
	b := Board{}
	b.Init()

	b_copy := b.DeepCopy()

	b_copy.board[f_a][r_1] = w_rook

	if b.board[f_a][r_1] != empty {
		t.Errorf("Original board modified. a1 is %s instead of %s", piece_to_string[b.board[f_a][r_1]], piece_to_string[empty])
	}
}
