package storage_repo

import (
	"context"
	"telegraph/models"
)

type StorageRepo interface {
	Init(ctx context.Context) error
	CreateAccount(account *models.Account) error
	UpdateAccountAccessToken(access_token string, new_access_token string) (*models.Account, error)
	GetAccountInfo(access_token string, fields []string) (*models.Account, error)
}
