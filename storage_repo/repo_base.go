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
	CreatePage(page *models.Page) error
	GetPage(path string) (*models.Page, error)
	EditPage(page_id int, page *models.Page) (*models.Page, error)
}
