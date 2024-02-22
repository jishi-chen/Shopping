package domain

import (
	"context"
)

// 定義屬性，當程式裡面創建時，必定這些屬性都要擁有。
type Account struct {
	ID     string `json:"id"`
	Name   string `json:"name" example:"account name"`
	Status string `json:"status"`
}

// 定義了 repository 層的各種方法，必須要按照這個定義實作，且呼叫的程式只能呼叫這些定義好的方法，不然會error。
type AccountRepository interface {
	GetByID(ctx context.Context, id string) (*Account, error)
	Store(ctx context.Context, d *Account) error
	UpdateStatus(ctx context.Context, d *Account) error
}

// 定義 usecase 層的各種方法
type AccountUsecase interface {
	GetByID(ctx context.Context, id string) (*Account, error)
	Store(ctx context.Context, d *Account) error
	UpdateStatus(ctx context.Context, d *Account) error
}
