package api_handler

import (
	"context"
	"crypto/md5"
	"fmt"
	"regexp"
	"telegraph/config"
	"telegraph/models"
	"telegraph/storage_repo"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ok     bool   `json:"ok"`
	Result any    `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
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

func GetPageUrl(path string) string {
	return fmt.Sprintf("http://%s:%d/getPage/%s", config.HOST, config.PORT, path)
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

func ReturnSuccess(c *gin.Context, status_code int, result any) {
	resp := Response{}
	resp.Ok = true
	resp.Result = result
	c.JSON(status_code, resp)
}

func ReturnError(c *gin.Context, status_code int, error_msg string) {
	resp := Response{}
	resp.Ok = false
	resp.Error = error_msg
	println(error_msg)
	c.JSON(status_code, resp)
}
