package repository

import (
	"context"

	"github.com/swopz/proof-feed/domain"
)

var NewPromoRepository = func() PromoRepository {
	return nil
}

type PromoRepository interface {
	Store(ctx context.Context, promo *domain.Promo) (int64, error)
	Load(ctx context.Context, promoID int64) (*domain.Promo, error)
}
