package storage_repo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"telegraph/config"
	"telegraph/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StorageRepoMongo struct {
	client *mongo.Client
	ctx    context.Context
}

func (s *StorageRepoMongo) Init(ctx context.Context) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetConfigs().DbUrl))
	if err != nil {
		msg := fmt.Sprintf("Failed opening connection to sqlite: %v", err)
		log.Println(msg)
		panic(err)
	}
	s.client = client
	s.ctx = ctx
	return nil
}

func (s *StorageRepoMongo) CreateAccount(account *models.Account) (*models.Account, error) {
	// TODO
	return nil, errors.New("Not implemented")
}

func (s *StorageRepoMongo) UpdateAccountInfo(access_token string, account *models.Account) (*models.Account, error) {
	// TODO
	return nil, errors.New("Not implemented")
}

func (s *StorageRepoMongo) UpdateAccountAccessToken(access_token string, new_access_token string) (*models.Account, error) {
	// TODO
	return nil, errors.New("Not implemented")
}

func (s *StorageRepoMongo) GetAccountInfo(access_token string, fields []string) (*models.Account, error) {
	// TODO
	return nil, errors.New("Not implemented")
}

func (s *StorageRepoMongo) CreatePage(page *models.Page) (*models.Page, error) {
	// TODO
	return nil, errors.New("Not implemented")
}

func (s *StorageRepoMongo) GetPage(path string) (*models.Page, error) {
	// TODO
	return nil, errors.New("Not implemented")
}

func (s *StorageRepoMongo) EditPage(page_id int, page *models.Page) (*models.Page, error) {
	// TODO
	return nil, errors.New("Not implemented")
}

func (s *StorageRepoMongo) ListPages(account_id int, limit int, offset int) (*models.PageList, error) {
	// TODO
	return nil, errors.New("Not implemented")
}
