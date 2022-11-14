package storage_repo

import (
	"context"
	"fmt"
	"log"
	"telegraph/models"
	"telegraph/storage_repo/ent"
	account_lib "telegraph/storage_repo/ent/account"
	page_lib "telegraph/storage_repo/ent/page"

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

func convert_account_type(u *ent.Account) *models.Account {
	fmt.Println(u)
	account := &models.Account{}
	account.Id = u.ID
	account.AuthorName = u.AuthorName
	account.ShortName = u.ShortName
	account.AccessToken = u.AccessToken
	account.AuthorUrl = u.AuthorURL
	account.AuthUrl = u.AuthURL
	fmt.Println(account)
	return account
}

func convert_page_type(u *ent.Page) *models.Page {
	page := &models.Page{}
	page.Id = u.ID
	page.AccountId = u.AccountID
	page.Content = u.Content
	page.Description = u.Description
	page.Path = u.Path
	page.Title = u.Title
	page.ImageUrl = u.ImageURL
	page.Url = u.URL
	page.AuthorUrl = u.AuthorURL
	page.AuthorName = u.AuthorName
	return page
}

func (s *StorageRepoSqlite3) CreateAccount(account *models.Account) (*models.Account, error) {
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
		return nil, err
	}
	log.Println("user was created: ", u)
	return convert_account_type(u), nil
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
	return convert_account_type(t), err
}

func (s *StorageRepoSqlite3) GetAccountInfo(access_token string, fields []string) (*models.Account, error) {
	t := s.client.Account.
		Query().
		Where(account_lib.AccessToken(access_token))
	var u *ent.Account = nil
	var err error = nil
	if len(fields) > 0 {
		u, err = t.Select(fields...).Only(s.ctx)
	} else {
		u, err = t.Only(s.ctx)
	}
	if err != nil {
		return nil, err
	}
	log.Println("user returned: ", u)
	account := convert_account_type(u)
	return account, nil
}

func (s *StorageRepoSqlite3) CreatePage(page *models.Page) (*models.Page, error) {
	p, err := s.client.Page.
		Create().
		SetAccountID(page.AccountId).
		SetTitle(page.Title).
		SetPath(page.Path).
		SetDescription(page.Description).
		SetImageURL(page.ImageUrl).
		SetAuthorName(page.AuthorName).
		SetAuthorURL(page.AuthorUrl).
		SetContent(page.Content).
		Save(s.ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("page was created: ", p)
	return convert_page_type(p), nil
}

func (s *StorageRepoSqlite3) GetPage(path string) (*models.Page, error) {
	p, err := s.client.Page.
		Query().
		Where(page_lib.Path(path)).
		Only(s.ctx)
	if err != nil {
		return nil, err
	}
	log.Println("page returned: ", p)
	page := convert_page_type(p)
	return page, nil
}

func (s *StorageRepoSqlite3) EditPage(page_id int, page *models.Page) (*models.Page, error) {
	t := s.client.Page.
		UpdateOneID(page_id)
	if page.Title != "" {
		t = t.SetTitle(page.Title)
	}
	if page.Description != "" {
		t = t.SetDescription(page.Description)
	}
	if page.ImageUrl != "" {
		t = t.SetImageURL(page.ImageUrl)
	}
	if page.AuthorName != "" {
		t = t.SetAuthorName(page.AuthorName)
	}
	if page.AuthorUrl != "" {
		t = t.SetAuthorURL(page.AuthorUrl)
	}
	if page.Content != "" {
		t = t.SetContent(page.Content)
	}
	u, err := t.Save(s.ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("page was updated: ", u)
	return convert_page_type(u), nil
}

func (s *StorageRepoSqlite3) ListPages(account_id int, limit int, offset int) (*models.PageList, error) {
	total_count, err := s.client.Page.
		Query().
		Where(page_lib.AccountID(account_id)).
		Count(s.ctx)
	if err != nil {
		return nil, err
	}
	pages_ent, err := s.client.Page.
		Query().
		Where(page_lib.AccountID(account_id)).
		Offset(offset).
		Limit(limit).
		All(s.ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("pages count: %d, offset: %d, limit: %d\n",
		len(pages_ent), offset, limit)
	pages := make([]*models.Page, 0, len(pages_ent))
	for i := 0; i < len(pages_ent); i++ {
		pages = append(pages, convert_page_type(pages_ent[i]))
	}
	page_list := &models.PageList{
		TotalCount: total_count,
		Pages:      pages,
		Offset:     offset,
		Limit:      limit,
		Count:      len(pages),
	}
	return page_list, nil
}
