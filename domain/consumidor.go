package domain

import "time"

type Consumidor struct {
	ID          int64
	Name        string
	Nasc        time.Time
	Preferences []*Segment
}
