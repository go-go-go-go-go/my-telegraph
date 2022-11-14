package api_handler

import (
	"context"
	"crypto/md5"
	"fmt"
	"regexp"
	"telegraph/models"
	"telegraph/storage_repo"
	"time"
)

type Response struct {
	Ok     bool   `json:"ok"`
	Result any    `json:"result"`
	Error  string `json:"error"`
}

func GenerateAccessToken() string {
	src := fmt.Sprint(time.Now().UnixNano())
	srcCode := md5.Sum([]byte(src))
	code := fmt.Sprintf("%x", srcCode)
	return code
}

func GeneratePagePath(title string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	f_title := reg.ReplaceAllString(title, "")
	path := fmt.Sprintf("%s-%s-%d", f_title,
		time.Now().Format("2006-01-02"), time.Now().Unix())
	return path
}

func ValidateAccessToken(token string) (*models.Account, error) {
	repo := storage_repo.GetStorageRepo(context.Background())
	account, err := repo.GetAccountInfo(token, []string{"id"})
	if err != nil {
		return nil, err
	} else {
		return account, nil
	}
}
