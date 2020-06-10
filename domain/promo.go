package domain

import "time"

const (
	OrderByDiscount    = PromoOrder("discount")
	OrderByPreferences = PromoOrder("preferences")
	OrderByCreation    = PromoOrder("creation")
)

type PromoOrder string
type Segment string

type Promo struct {
	ID         int64
	Name       string
	DtCreation time.Time
	Value      float64
	Discount   float64
	Segment    Segment
}
