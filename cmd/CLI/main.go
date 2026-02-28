package main

import (
	"fmt"
	"time"
)

type BIN struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []BIN
}

func createBIN(id string, name string, private bool) BIN {
	return BIN{
		id:        id,
		name:      name,
		private:   private,
		createdAt: time.Now(),
	}
}

func NewBinList() *BinList {
	return &BinList{
		bins: make([]BIN, 0),
	}
}

func (bl *BinList) Add(bin BIN) {
	bl.bins = append(bl.bins, bin)
}

func main() {

}