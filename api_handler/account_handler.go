package api_handler

import (
	"context"
	"crypto/md5"
	"fmt"
	"net/http"
	"strings"
	"telegraph/models"
	"telegraph/storage_repo"
	"time"

	"github.com/gin-gonic/gin"
)

type account_response struct {
	Ok     bool            `json:"ok"`
	Result *models.Account `json:"result"`
	Error  string          `json:"error"`
}

func CreateAccount(c *gin.Context) {
	var account models.Account
	err := c.ShouldBindQuery(&account)
	if err == nil {
		fmt.Println("Prepared to create new account", account)
		resp := create(&account)
		c.JSON(http.StatusOK, resp)
	} else {
		msg := fmt.Sprintf("Failed to parse query: %s", err)
		fmt.Println(msg)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": msg,
		})
	}
}

func GenerateAccessToken() string {
	src := fmt.Sprint(time.Now().UnixNano())
	srcCode := md5.Sum([]byte(src))
	code := fmt.Sprintf("%x", srcCode)
	return code
}

func create(account *models.Account) *account_response {
	account.AccessToken = GenerateAccessToken()

	repo := storage_repo.GetStorageRepo(context.Background())
	err := repo.CreateAccount(account)

	resp := account_response{}
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
	}
	s := make([]string, 0)
	if fields != "" {
		t := strings.Split(strings.Trim(fields, " []"), ",")
		for i := 0; i < len(t); i++ {
			s = append(s, strings.Trim(t[i], " \"'"))
		}
	}
	resp := fetch(access_token, s)
	c.JSON(http.StatusOK, resp)
}

func fetch(access_token string, fields []string) *account_response {

	repo := storage_repo.GetStorageRepo(context.Background())
	account, err := repo.GetAccountInfo(access_token, fields)

	resp := account_response{}
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
	}
	new_access_token := GenerateAccessToken()

	repo := storage_repo.GetStorageRepo(context.Background())
	account, err := repo.UpdateAccountAccessToken(access_token, new_access_token)

	resp := account_response{}
	if err != nil {
		resp.Ok = false
		resp.Error = fmt.Sprint(err)
	} else {
		resp.Ok = true
		resp.Result = account
	}
	c.JSON(http.StatusOK, resp)
}
