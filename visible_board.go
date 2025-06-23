package main

type VisibleSquare struct {
	sq_status string // can be neutral, selected, target
	sq_type   string
	sq_x      float64
	sq_y      float64
}

type VisibleBoard struct {
	Squares      [8][8]VisibleSquare
	SelectedFile Col
	SelectedRank Rank
}

func (b *VisibleBoard) Init() {
	for rank_index := r_8; rank_index >= r_1; rank_index-- {
		for file_index := f_a; file_index <= f_h; file_index++ {
			sq_type := "tile1"
			even_row_odd_column := file_index%2 == 0 && rank_index%2 != 0
			odd_row_event_column := file_index%2 != 0 && rank_index%2 == 0
			if even_row_odd_column || odd_row_event_column {
				sq_type = "tile2"
			}
			b.Squares[file_index][rank_index] = VisibleSquare{
				sq_status: "neutral",
				sq_type:   sq_type,
				sq_x:      float64(int(file_index) * tile_size),
				sq_y:      float64((7 - int(rank_index)) * tile_size),
			}
		}
	}
}

func (b *VisibleBoard) ResetAllTiles() {
	for rank_index := r_8; rank_index >= r_1; rank_index-- {
		for file_index := f_a; file_index <= f_h; file_index++ {
			b.Squares[file_index][rank_index].sq_status = "neutral"
		}
	}
	b.SelectedFile = f_none
	b.SelectedRank = r_none
}

func (b *VisibleBoard) SelectTile(file Col, rank Rank) {
	b.ResetAllTiles()
	b.SelectedFile = file
	b.SelectedRank = rank
	b.Squares[b.SelectedFile][b.SelectedRank].sq_status = "selected"
}

func (b *VisibleBoard) TargetTile(file Col, rank Rank) {
	b.Squares[file][rank].sq_status = "target"
}
