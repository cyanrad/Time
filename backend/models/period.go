package models

import (
	"fmt"
)

type ActivityType string

var types = []ActivityType{"Work", "Positive", "People", "Neutral", "University", "Sleep", "Rest", "Useless", "Mandatory", "Unknown"}

type Period struct {
	ID          int64
	DateStamp   int32 `json:"date_stamp" db:"date_stamp"`
	Location    int16
	Duration    int16
	Type        ActivityType
	Title       string
	Description string
}

func (p *Period) Print() {
	fmt.Printf("%#v\n", p)
}
