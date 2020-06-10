package proxy

import (
	"context"

	"github.com/swopz/proof-feed/domain"
	"github.com/swopz/proof-feed/service"
)

var NewFeedProxy = func() service.FeedService {
	return &feedProxy{}
}

type feedProxy struct{}

// Create gera um novo feed para o consumidor com as promoções selecionadas.
func (proxy *feedProxy) Create(ctx context.Context, consumidor int64, promos []int64) error {
	return nil
}

// Improve adiciona novas promoções ao feed.
func (proxy *feedProxy) Improve(ctx context.Context, consumidor int64, promos []int64) error {
	return nil
}

// Invalidate invalida uma promoção removendo-a de todos os feeds gerados.
func (proxy *feedProxy) Invalidate(ctx context.Context, promoID int64) error {
	return nil
}

// Load carrega uma página do feed para o consumidor.
func (proxy *feedProxy) Load(ctx context.Context, consumidorID int64, pageNo int, order domain.PromoOrder) ([]*domain.Promo, error) {
	return nil, nil
}
