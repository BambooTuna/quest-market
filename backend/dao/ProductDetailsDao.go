package dao

import (
	"context"
	"github.com/BambooTuna/quest-market/backend/model/goods"
)

type ProductDetailsDao interface {
	ResolveByProductId(ctx context.Context, productId string, practitionerId string) (*goods.ProductDetails, error)
	ResolveByPresenterId(ctx context.Context, practitionerId string) ([]goods.ProductDetails, error)
	Insert(ctx context.Context, record *goods.ProductDetails) error
	Update(ctx context.Context, record *goods.ProductDetails) (int64, error)
}
