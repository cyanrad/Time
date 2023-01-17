package models

import "fmt"

type Achievement struct {
	ID        int64
	DateStamp int32 `json:"date_stamp" db:"date_stamp"`
	Text      string
}

func (a *Achievement) Print() {
	fmt.Printf("%#v\n", a)
}
