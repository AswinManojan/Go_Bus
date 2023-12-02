package entities

import (
	"encoding/json"

	"gorm.io/gorm"
)

type SeatLayoutStr struct {
	DB *gorm.DB
}

func NewSeatLAyout(db *gorm.DB) *SeatLayoutStr {
	return &SeatLayoutStr{
		DB: db,
	}
}

type BusSeatLayout struct {
	gorm.Model
	// SeatLayoutId   int    `json:"seat_id"`
	DeckOneColumns int    `json:"deck_one_columns"`
	DeckTwoColumns int    `json:"deck_two_columns"`
	DeckOneRows    int    `json:"deck_one_rows"`
	DeckTwoRows    int    `json:"deck_two_rows"`
	DeckOneLayout  []byte `json:"deckone_seat_layout"`
	DeckTwoLayout  []byte `json:"decktwo_seat_layout"`
}

type DeckOneLayoutstr struct {
	DeckOneLayout [][]bool
}
type DeckTwoLayoutstr struct {
	DeckTwoLayout [][]bool
}

func (slr *SeatLayoutStr) Layout1() {
	layout := &BusSeatLayout{}
	layout.DeckOneColumns = 3
	layout.DeckOneRows = 4
	layout.DeckTwoColumns = 3
	layout.DeckTwoRows = 4
	d1 := &DeckOneLayoutstr{}
	d2 := &DeckTwoLayoutstr{}
	for i := 1; i <= layout.DeckOneRows; i++ {
		newRow := []bool{}
		for j := 1; j <= layout.DeckOneColumns; j++ {
			newRow = append(newRow, false)
		}
		d1.DeckOneLayout = append(d1.DeckOneLayout, newRow)
	}
	for i := 1; i <= layout.DeckTwoRows; i++ {
		newRow := []bool{}
		for j := 1; j <= layout.DeckTwoColumns; j++ {
			newRow = append(newRow, false)
		}
		d2.DeckTwoLayout = append(d2.DeckTwoLayout, newRow)
	}
	d1layout, _ := json.Marshal(&d1)
	d2layout, _ := json.Marshal(&d2)
	layout.DeckOneLayout = d1layout
	layout.DeckTwoLayout = d2layout

	slr.DB.Create(&layout)
}

func (slr *SeatLayoutStr) Layout2() {
	layout := &BusSeatLayout{}
	layout.DeckOneColumns = 3
	layout.DeckOneRows = 10
	layout.DeckTwoColumns = 3
	layout.DeckTwoRows = 4
	d1 := &DeckOneLayoutstr{}
	d2 := &DeckTwoLayoutstr{}
	for i := 1; i <= layout.DeckOneRows; i++ {
		newRow := []bool{}
		for j := 1; j <= layout.DeckOneColumns; j++ {
			newRow = append(newRow, false)
		}
		d1.DeckOneLayout = append(d1.DeckOneLayout, newRow)
	}
	for i := 1; i <= layout.DeckTwoRows; i++ {
		newRow := []bool{}
		for j := 1; j <= layout.DeckTwoColumns; j++ {
			newRow = append(newRow, false)
		}
		d2.DeckTwoLayout = append(d2.DeckTwoLayout, newRow)
	}
	d1layout, _ := json.Marshal(&d1)
	d2layout, _ := json.Marshal(&d2)
	layout.DeckOneLayout = d1layout
	layout.DeckTwoLayout = d2layout
	slr.DB.Create(&layout)

}
func (slr *SeatLayoutStr) Layout3() {
	layout := &BusSeatLayout{}
	layout.DeckOneColumns = 3
	layout.DeckOneRows = 10
	layout.DeckTwoColumns = 0
	layout.DeckTwoRows = 0
	d1 := &DeckOneLayoutstr{}
	d2 := &DeckTwoLayoutstr{}
	for i := 1; i <= layout.DeckOneRows; i++ {
		newRow := []bool{}
		for j := 1; j <= layout.DeckOneColumns; j++ {
			newRow = append(newRow, false)
		}
		d1.DeckOneLayout = append(d1.DeckOneLayout, newRow)
	}
	for i := 1; i <= layout.DeckTwoRows; i++ {
		newRow := []bool{}
		for j := 1; j <= layout.DeckTwoColumns; j++ {
			newRow = append(newRow, false)
		}
		d2.DeckTwoLayout = append(d2.DeckTwoLayout, newRow)
	}
	d1layout, _ := json.Marshal(&d1)
	d2layout, _ := json.Marshal(&d2)
	layout.DeckOneLayout = d1layout
	layout.DeckTwoLayout = d2layout
	slr.DB.Create(&layout)
}
