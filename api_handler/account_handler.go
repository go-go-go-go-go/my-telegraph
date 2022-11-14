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
		resp := create(&account)
		c.JSON(http.StatusOK, resp)
		return
	} else {
		msg := fmt.Sprintf("Failed to parse query: %s", err)
		fmt.Println(msg)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": msg,
		})
		return
	}
}

func create(account *models.Account) *Response {
	account.AccessToken = GenerateAccessToken()

	repo := storage_repo.GetStorageRepo(context.Background())
	err := repo.CreateAccount(account)

	resp := Response{}
	if err != nil {
		resp.Ok = false
		resp.Error = fmt.Sprint(err)
	} else {
		resp.Ok = true
		resp.Result = account
	}
	return &resp
}

func GetAccountInfo(c *gin.Context) {
	access_token := c.Query("access_token")
	fields := c.DefaultQuery("fields", "")
	if access_token == "" {
		msg := "access_token is required"
		fmt.Println(msg)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": msg,
		})
		return
	}
	s := make([]string, 0)
	if fields != "" {
		t := strings.Split(strings.Trim(fields, " []"), ",")
		for i := 0; i < len(t); i++ {
			s = append(s, strings.Trim(t[i], " \"'"))
		}
	}
	resp := fetch(access_token, s)
	if resp.Ok {
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusBadRequest, resp)
	}
}

func fetch(access_token string, fields []string) *Response {

	repo := storage_repo.GetStorageRepo(context.Background())
	account, err := repo.GetAccountInfo(access_token, fields)

	resp := Response{}
	if err != nil {
		resp.Ok = false
		resp.Error = fmt.Sprint(err)
	} else {
		resp.Ok = true
		resp.Result = account
	}
	return &resp
}

func RevokeAccessToken(c *gin.Context) {
	access_token := c.Query("access_token")
	if access_token == "" {
		msg := "access_token is required"
		fmt.Println(msg)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": msg,
		})
		return
	}
	new_access_token := GenerateAccessToken()

	repo := storage_repo.GetStorageRepo(context.Background())
	account, err := repo.UpdateAccountAccessToken(access_token, new_access_token)

	resp := Response{}
	if err != nil {
		resp.Ok = false
		resp.Error = fmt.Sprint(err)
	} else {
		resp.Ok = true
		resp.Result = account
	}
	c.JSON(http.StatusOK, resp)
}
