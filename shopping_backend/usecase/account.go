package usecase

import (
	"context"
	"errors"
	"shopping_backend/domain"

	uuid "github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// 定義依賴注入(DI)所需的Repository層: 注入了之後我們再對 Repository 層做各種邏輯的操作，要注意的是，雖然目前只有注入所屬 PostgreSQL 的 repository，但如果 account 有很多不同的來源，比如說 MongoDB、Microservice 等等，我們可以注入更多 repository 來操作。
type accountUsecase struct {
	accountRepo domain.AccountRepository
}

// 設計確實注入的function: 與 repository 層一樣，需要一個 function 來要求要注入哪些 repository。
func NewAccountUsecase(accountRepo domain.AccountRepository) domain.AccountUsecase {
	return &accountUsecase{
		accountRepo: accountRepo,
	}
}

// 依照domain.AccountRepository interface來實作: 與 repository 層一樣，要以 interface 規範來實作，不過這裡是實作業務邏輯。
func (du *accountUsecase) GetByID(ctx context.Context, id string) (*domain.Account, error) {
	aAccount, err := du.accountRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return aAccount, nil
}

func (du *accountUsecase) Store(ctx context.Context, d *domain.Account) error {
	if d.ID == "" {
		d.ID = uuid.Must(uuid.NewV4()).String()
	}
	if d.Status == "" {
		d.Status = "good"
	}
	if err := du.accountRepo.Store(ctx, d); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (du *accountUsecase) UpdateStatus(ctx context.Context, d *domain.Account) error {
	if d.Status == "" {
		err := errors.New("Status is blank")
		logrus.Error(err)
		return err
	}

	if err := du.accountRepo.UpdateStatus(ctx, d); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
