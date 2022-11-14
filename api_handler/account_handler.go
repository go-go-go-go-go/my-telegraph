package api_handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"telegraph/models"
	"telegraph/storage_repo"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var account models.Account
	err := c.ShouldBindQuery(&account)
	if err == nil {
		fmt.Println("Prepared to create new account", account)
		a, err := create_account(&account)
		if err == nil {
			ReturnSuccess(c, http.StatusOK, a)
		} else {
			ReturnError(c, http.StatusBadRequest, err.Error())
		}

		return
	} else {
		msg := fmt.Sprintf("Failed to parse query: %s", err)
		ReturnError(c, http.StatusBadRequest, msg)
		return
	}
}

func create_account(account *models.Account) (*models.Account, error) {
	account.AccessToken = GenerateAccessToken()
	repo := storage_repo.GetStorageRepo(context.Background())
	a, err := repo.CreateAccount(account)

	if err != nil {
		return nil, err
	} else {
		return a, nil
	}
}

func GetAccountInfo(c *gin.Context) {
	access_token := c.Query("access_token")
	fields := c.DefaultQuery("fields", "")
	if access_token == "" {
		msg := "access_token is required"
		ReturnError(c, http.StatusBadRequest, msg)
		return
	}
	s := make([]string, 0)
	if fields != "" {
		t := strings.Split(strings.Trim(fields, " []"), ",")
		for i := 0; i < len(t); i++ {
			s = append(s, strings.Trim(t[i], " \"'"))
		}
	}
	a, err := fetch_account_info(access_token, s)
	if err == nil {
		ReturnSuccess(c, http.StatusOK, a)
	} else {
		ReturnError(c, http.StatusBadRequest, err.Error())
	}
}

func fetch_account_info(access_token string, fields []string) (*models.Account, error) {

	repo := storage_repo.GetStorageRepo(context.Background())
	account, err := repo.GetAccountInfo(access_token, fields)

	if err != nil {
		return nil, err
	}
	page_count, err := GetPageCount(access_token)
	if err != nil {
		return nil, err
	}
	account.PageCount = page_count
	return account, nil
}

func RevokeAccessToken(c *gin.Context) {
	access_token := c.Query("access_token")
	if access_token == "" {
		msg := "access_token is required"
		ReturnError(c, http.StatusBadRequest, msg)
		return
	}
	new_access_token := GenerateAccessToken()

	repo := storage_repo.GetStorageRepo(context.Background())
	account, err := repo.UpdateAccountAccessToken(access_token, new_access_token)

	if err != nil {
		ReturnError(c, http.StatusBadRequest, err.Error())
	} else {
		ReturnSuccess(c, http.StatusOK, account)
	}
}

func EditAccountInfo(c *gin.Context) {
	var account models.Account
	err := c.ShouldBindQuery(&account)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, err.Error())
	}

	repo := storage_repo.GetStorageRepo(context.Background())
	a, err := repo.UpdateAccountInfo(account.AccessToken, &account)

	if err != nil {
		ReturnError(c, http.StatusBadRequest, err.Error())
	} else {
		ReturnSuccess(c, http.StatusOK, a)
	}
}
