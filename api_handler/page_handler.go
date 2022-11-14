package api_handler

import (
	"context"
	"fmt"
	"net/http"
	"telegraph/config"
	"telegraph/models"
	"telegraph/storage_repo"

	"github.com/gin-gonic/gin"
)

func CreatePage(c *gin.Context) {
	var page models.Page
	err := c.ShouldBindQuery(&page)
	if err != nil {
		msg := fmt.Sprintf("Failed to parse query: %s", err)
		fmt.Println(msg)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": msg,
		})
		return
	}
	account, err := ValidateAccessToken(page.AccessToken)
	if err != nil {
		msg := fmt.Sprintf("Failed to validate access token: %s", err)
		fmt.Println(msg)
		c.JSON(http.StatusForbidden, gin.H{
			"error": msg,
		})
		return
	}
	println(account)
	page.AccountId = account.Id
	fmt.Println("Prepared to create new page", page)
	resp := createPage(&page)
	if resp.Ok {
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusBadRequest, resp)
	}
}

func createPage(page *models.Page) *Response {
	page.Path = GeneratePagePath(page.Title)
	repo := storage_repo.GetStorageRepo(context.Background())
	err := repo.CreatePage(page)

	resp := Response{}
	if err != nil {
		resp.Ok = false
		resp.Error = fmt.Sprint(err)
	} else {
		resp.Ok = true
		if !page.ReturnContent {
			page.Content = ""
		}
		resp.Result = page
	}
	return &resp
}

func GetPage(c *gin.Context) {
	path := c.Param("path")
	return_content_str := c.DefaultQuery("return_content", "false")
	return_content := true
	if return_content_str == "true" {
		return_content = true
	} else if return_content_str == "false" {
		return_content = false
	} else {
		msg := "access_token is required"
		fmt.Println(msg)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": msg,
		})
		return
	}
	resp := fetch_page(path, return_content)
	if resp.Ok {
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusBadRequest, resp)
	}
}

func fetch_page(path string, return_content bool) *Response {
	repo := storage_repo.GetStorageRepo(context.Background())
	page, err := repo.GetPage(path)

	resp := Response{}
	if err != nil {
		resp.Ok = false
		resp.Error = fmt.Sprint(err)
	} else {
		resp.Ok = true
		page.Url = fmt.Sprintf("http://%s:%d/getPage/%s", config.HOST, config.PORT, path)
		if !return_content {
			page.Content = ""
		}
		resp.Result = page
	}
	return &resp
}

func EditPage(c *gin.Context) {
	path := c.Param("path")
	var page models.Page
	err := c.ShouldBindQuery(&page)
	if err != nil {
		msg := fmt.Sprintf("Failed to parse query: %s", err)
		fmt.Println(msg)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": msg,
		})
		return
	}
	account, err := ValidateAccessToken(page.AccessToken)
	if err != nil {
		msg := fmt.Sprintf("Failed to validate access token: %s", err)
		fmt.Println(msg)
		c.JSON(http.StatusForbidden, gin.H{
			"error": msg,
		})
		return
	}
	// println(account)
	repo := storage_repo.GetStorageRepo(context.Background())
	current_page, err := repo.GetPage(path)
	if err != nil {
		msg := fmt.Sprintf("Failed to get page by path: %s", path)
		fmt.Println(msg)
		c.JSON(http.StatusNotFound, gin.H{
			"error": msg,
		})
		return
	}
	if current_page.AccountId != account.Id {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Access token is wrong",
		})
		return
	}
	page.Path = path
	page.Id = current_page.Id
	fmt.Println("Prepared to update page", page)
	resp := editPage(&page)
	if resp.Ok {
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusBadRequest, resp)
	}
}

func editPage(page *models.Page) *Response {
	return_content := page.ReturnContent
	page.Path = GeneratePagePath(page.Title)
	repo := storage_repo.GetStorageRepo(context.Background())
	page, err := repo.EditPage(page.Id, page)
	resp := Response{}
	if err != nil {
		resp.Ok = false
		resp.Error = fmt.Sprint(err)
	} else {
		resp.Ok = true
		if !return_content {
			page.Content = ""
		}
		resp.Result = page
	}
	return &resp
}
