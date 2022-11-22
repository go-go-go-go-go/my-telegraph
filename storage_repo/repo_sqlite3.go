package storage_repo

import (
	"context"
	"log"
	"telegraph/models"
	"telegraph/storage_repo/ent"
	account_lib "telegraph/storage_repo/ent/account"

	_ "github.com/mattn/go-sqlite3"
)

type StorageRepoSqlite3 struct {
	client *ent.Client
	ctx    context.Context
}

func (s *StorageRepoSqlite3) Init(ctx context.Context) error {
	client, err := ent.Open("sqlite3", "file:data/telegraph.db?_fk=1")
	if err != nil {
		log.Printf("failed opening connection to sqlite: %v", err)
		return err
	}
	// defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Printf("failed creating schema resources: %v", err)
		return err
	}
	s.client = client
	s.ctx = ctx
	return nil
}

func convert_type(u *ent.Account) *models.Account {
	account := &models.Account{}
	account.ID = u.ID
	account.AuthorName = u.AuthorName
	account.ShortName = u.ShortName
	account.AccessToken = u.AccessToken
	account.AuthorUrl = u.AuthorURL
	account.AuthUrl = u.AuthURL
	return account
}

func (s *StorageRepoSqlite3) CreateAccount(account *models.Account) error {
	u, err := s.client.Account.
		Create().
		SetShortName(account.ShortName).
		SetAuthorName(account.AuthorName).
		SetAuthorURL(account.AuthorUrl).
		SetAccessToken(account.AccessToken).
		SetAuthURL(account.AuthUrl).
		Save(s.ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("user was created: ", u)
	return nil
}

func (s *StorageRepoSqlite3) UpdateAccountAccessToken(access_token string, new_access_token string) (*models.Account, error) {
	u, err := s.client.Account.
		Query().
		Where(account_lib.AccessToken(access_token)).
		Only(s.ctx)
	if err != nil {
		return nil, err
	}
	t, err := u.Update().SetAccessToken(new_access_token).Save(s.ctx)
	if err != nil {
		return nil, err
	}
	return convert_type(t), err
}

func (s *StorageRepoSqlite3) GetAccountInfo(access_token string, fields []string) (*models.Account, error) {
	u, err := s.client.Account.
		Query().
		Where(account_lib.AccessToken(access_token)).
		Select(fields...).
		Only(s.ctx)
	if err != nil {
		return nil, err
	}
	log.Println("user returned: ", u)
	account := convert_type(u)
	return account, nil
}
