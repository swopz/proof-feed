package service

import (
	"context"

	"github.com/swopz/proof-feed/domain"
)

// FeedService é a interface disponibilizada pelo microserviço.
type FeedService interface {
	// Create gera um novo feed para o consumidor com as promoções selecionadas.
	Create(ctx context.Context, consumidor int64, promos []int64) error
	// Improve adiciona novas promoções ao feed.
	Improve(ctx context.Context, consumidor int64, promos []int64) error
	// Invalidate invalida uma promoção removendo-a de todos os feeds gerados.
	Invalidate(ctx context.Context, promoID int64) error
	// Load carrega uma página do feed para o consumidor.
	Load(ctx context.Context, consumidorID int64, pageNo int, order domain.PromoOrder) ([]*domain.Promo, error)
}

type GeneratorService interface {
	CreateConsumidores(ctx context.Context, consumidores []*domain.Consumidor) error
	CreatePromos(ctx context.Context, promos []*domain.Promo) error
}
