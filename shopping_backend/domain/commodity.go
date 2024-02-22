package domain

import "context"

type Commodity struct {
	ID     string
	UserID string
	Name   string
}

type CommodityRepository interface {
	GetByID(ctx context.Context, id string) (*Commodity, error)
	Store(ctx context.Context, d *Commodity) error
}

type CommodityUsecase interface {
	GetByID(ctx context.Context, id string) (*Commodity, error)
	Store(ctx context.Context, d *Commodity) error
}
