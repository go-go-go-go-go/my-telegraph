package storage_repo

import (
	"context"
	"telegraph/models"
)

type StorageRepo interface {
	Init(ctx context.Context) error
	CreateAccount(account *models.Account) (*models.Account, error)
	UpdateAccountInfo(access_token string, account *models.Account) (*models.Account, error)
	UpdateAccountAccessToken(access_token string, new_access_token string) (*models.Account, error)
	GetAccountInfo(access_token string, fields []string) (*models.Account, error)
	CreatePage(page *models.Page) (*models.Page, error)
	GetPage(path string, view_add bool) (*models.Page, error)
	EditPage(page_id int, page *models.Page) (*models.Page, error)
	ListPages(account_id int, limit int, offset int) (*models.PageList, error)
	GetPageView(path string, year int, month int, day int, hour int) (int, error)
}
