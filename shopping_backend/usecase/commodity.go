package usecase

import (
	"context"

	"shopping_backend/domain"

	"github.com/sirupsen/logrus"
)

type commodityUsecase struct {
	commodityRepo domain.CommodityRepository
}

// NewCommodityUsecase ...
func NewCommodityUsecase(commodityRepo domain.CommodityRepository) domain.CommodityUsecase {
	return &commodityUsecase{
		commodityRepo,
	}
}

func (du *commodityUsecase) GetByID(ctx context.Context, id string) (*domain.Commodity, error) {
	aCommodity, err := du.commodityRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return aCommodity, nil
}

func (du *commodityUsecase) Store(ctx context.Context, d *domain.Commodity) error {
	if err := du.commodityRepo.Store(ctx, d); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
