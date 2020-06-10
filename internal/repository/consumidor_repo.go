package repository

import (
	"context"

	"github.com/swopz/proof-feed/domain"
)

var NewConsumidorRepository = func() ConsumidorRepository {
	return nil
}

type ConsumidorRepository interface {
	Store(ctx context.Context, consumidor *domain.Consumidor) (int64, error)
	Load(ctx context.Context, consumidorID int64) (*domain.Consumidor, error)
}
