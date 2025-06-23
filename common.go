package main

var assets_path = map[string]string{
	"tile1": "img/Tile2.png",
	"tile2": "img/Tile1.png",
	"b":     "img/black_bishop.png",
	"k":     "img/black_king.png",
	"n":     "img/black_knight.png",
	"p":     "img/black_pawn.png",
	"q":     "img/black_queen.png",
	"r":     "img/black_rook.png",
	"B":     "img/white_bishop.png",
	"K":     "img/white_king.png",
	"N":     "img/white_knight.png",
	"P":     "img/white_pawn.png",
	"Q":     "img/white_queen.png",
	"R":     "img/white_rook.png",
}

type Player int

const (
	// ranks list
	white Player = iota
	black
	p_none
)

var player_to_string = map[Player]string{
	white:  "white",
	black:  "black",
	p_none: "none",
}

type Rank int

const (
	// ranks list
	r_1 Rank = iota
	r_2
	r_3
	r_4
	r_5
	r_6
	r_7
	r_8
	r_none
)

var rank_to_string = map[Rank]string{
	r_1:    "1",
	r_2:    "2",
	r_3:    "3",
	r_4:    "4",
	r_5:    "5",
	r_6:    "6",
	r_7:    "7",
	r_8:    "8",
	r_none: "-",
}

// can't call the new type File
type Col int

const (
	// files list
	f_a Col = iota
	f_b
	f_c
	f_d
	f_e
	f_f
	f_g
	f_h
	f_none
)

var file_to_string = map[Col]string{
	f_a:    "a",
	f_b:    "b",
	f_c:    "c",
	f_d:    "d",
	f_e:    "e",
	f_f:    "f",
	f_g:    "g",
	f_h:    "h",
	f_none: "-",
}

type Piece int

const (
	// black pieces
	b_pawn Piece = iota
	b_rook
	b_knight
	b_bishop
	b_queen
	b_king
	// white pieces
	w_pawn
	w_rook
	w_knight
	w_bishop
	w_queen
	w_king
	// empty square
	empty
	off_board
)

type PieceType int

const (
	pawn PieceType = iota
	rook
	knight
	bishop
	queen
	king
	type_none
)

var piece_to_piece_type = map[Piece]PieceType{
	b_pawn:    pawn,
	b_rook:    rook,
	b_knight:  knight,
	b_bishop:  bishop,
	b_queen:   queen,
	b_king:    king,
	w_pawn:    pawn,
	w_rook:    rook,
	w_knight:  knight,
	w_bishop:  bishop,
	w_queen:   queen,
	w_king:    king,
	empty:     type_none,
	off_board: type_none,
}

var piece_to_string = map[Piece]string{
	b_pawn:    "p",
	b_rook:    "r",
	b_knight:  "n",
	b_bishop:  "b",
	b_queen:   "q",
	b_king:    "k",
	w_pawn:    "P",
	w_rook:    "R",
	w_knight:  "N",
	w_bishop:  "B",
	w_queen:   "Q",
	w_king:    "K",
	empty:     ".",
	off_board: "*",
}

var piece_to_player = map[Piece]Player{
	b_pawn:    black,
	b_rook:    black,
	b_knight:  black,
	b_bishop:  black,
	b_queen:   black,
	b_king:    black,
	w_pawn:    white,
	w_rook:    white,
	w_knight:  white,
	w_bishop:  white,
	w_queen:   white,
	w_king:    white,
	empty:     p_none,
	off_board: p_none,
}

type Move struct {
	start_file          Col
	start_rank          Rank
	end_file            Col
	end_rank            Rank
	piece               Piece
	turn_to_queen       bool
	is_capture          bool
	is_pawn_double_move bool
	is_long_castle      bool
	is_short_castle     bool
	is_en_passant       bool
}

func (m Move) Is(s_file Col, s_rank Rank, e_file Col, e_rank Rank) bool {
	start_is_same := m.start_file == s_file && m.start_rank == s_rank
	end_is_same := m.end_file == e_file && m.end_rank == e_rank
	return start_is_same && end_is_same
}

func (m Move) Equal(m2 Move) bool {
	start_is_same := m.start_file == m2.start_file && m.start_rank == m2.start_rank
	end_is_same := m.end_file == m2.end_file && m.end_rank == m2.end_rank
	return start_is_same && end_is_same
}

// this will not handle cases where
func (m Move) String() string {
	move := file_to_string[m.start_file]
	move += rank_to_string[m.start_rank]
	if m.is_capture {
		move += "x"
	}
	move += file_to_string[m.end_file]
	move += rank_to_string[m.end_rank]

	return move
}

func DeepCopy2DInt(src [][]Piece) [][]Piece {
	if src == nil {
		return nil
	}
	new_2d := make([][]Piece, len(src))
	for i := range src {
		new_2d[i] = make([]Piece, len(src[i]))
		copy(new_2d[i], src[i])
	}
	return new_2d
}

func DeepCopy8x8Piece(src [8][8]Piece) [8][8]Piece {
	var dst [8][8]Piece
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			dst[i][j] = src[i][j]
		}
	}
	return dst
}

type GameOptions struct {
	Mode string // pvp or pvai
}

// this struct is for the main menu of the game
// each labels will be displayed on the screen
// and the user can pick one of them
// if IsSelected is True then a marker will be displayed
type GameMenuItem struct {
	Label string
	X     float64
	Y     float64
}

type GameMenu struct {
	Items         []GameMenuItem
	SelectedIndex float64
}
