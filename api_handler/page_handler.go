package api_handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"telegraph/models"
	"telegraph/storage_repo"

	"github.com/gin-gonic/gin"
)

func CreatePage(c *gin.Context) {
	var page_req models.PageRequest
	err := c.ShouldBindQuery(&page_req)
	if err != nil {
		msg := fmt.Sprintf("Failed to parse query: %s", err)
		ReturnError(c, http.StatusBadRequest, msg)
		return
	}
	account, err := ValidateAccessToken(page_req.AccessToken)
	if err != nil {
		msg := fmt.Sprintf("Failed to validate access token: %s", err)
		ReturnError(c, http.StatusForbidden, msg)
		return
	}
	println(account)
	page, err := models.MakePage(&page_req)
	if err != nil {
		msg := fmt.Sprintf("Failed to parse content: %s", err)
		ReturnError(c, http.StatusForbidden, msg)
		return
	}
	page.AccountId = account.Id
	fmt.Println("Prepared to create new page", page)
	p, err := createPage(page)
	if err == nil {
		ReturnSuccess(c, http.StatusOK, p)
	} else {
		ReturnError(c, http.StatusBadRequest, err.Error())
	}
}

func createPage(page *models.Page) (*models.Page, error) {
	return_content := page.ReturnContent
	page.Path = GeneratePagePath(page.Title)
	repo := storage_repo.GetStorageRepo(context.Background())
	p, err := repo.CreatePage(page)

	if err != nil {
		return nil, err
	} else {
		if !return_content {
			p.Content = nil
		}
		p.CanEdit = true
		p.Url = GetPageUrl(p.Path)
		return p, nil
	}
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
		ReturnError(c, http.StatusBadRequest, msg)
		return
	}
	p, err := fetch_page(path, return_content)
	if err == nil {
		ReturnSuccess(c, http.StatusOK, p)
	} else {
		ReturnError(c, http.StatusBadRequest, err.Error())
	}
}

func fetch_page(path string, return_content bool) (*models.Page, error) {
	repo := storage_repo.GetStorageRepo(context.Background())
	page, err := repo.GetPage(path)

	if err != nil {
		return nil, err
	} else {
		page.Url = GetPageUrl(path)
		page_views, err := repo.GetPageView(path, -1, -1, -1, -1)
		if err != nil {
			page_views = 0
		}
		page.Views = page_views
		if !return_content {
			page.Content = nil
		}
		page.CanEdit = false
		return page, nil
	}
}

func EditPage(c *gin.Context) {
	path := c.Param("path")
	var page models.Page
	err := c.ShouldBindQuery(&page)
	if err != nil {
		msg := fmt.Sprintf("Failed to parse query: %s", err)
		ReturnError(c, http.StatusBadRequest, msg)
		return
	}
	account, err := ValidateAccessToken(page.AccessToken)
	if err != nil {
		msg := fmt.Sprintf("Failed to validate access token: %s", err)
		ReturnError(c, http.StatusForbidden, msg)
		return
	}
	// println(account)
	repo := storage_repo.GetStorageRepo(context.Background())
	current_page, err := repo.GetPage(path)
	if err != nil {
		msg := fmt.Sprintf("Failed to get page by path: %s", path)
		ReturnError(c, http.StatusNotFound, msg)
		return
	}
	if current_page.AccountId != account.Id {
		ReturnError(c, http.StatusForbidden, "Access token is wrong")
		return
	}
	page.Path = path
	page.Id = current_page.Id
	fmt.Println("Prepared to update page", page)
	p, err := editPage(&page)
	if err == nil {
		ReturnSuccess(c, http.StatusOK, p)
	} else {
		ReturnError(c, http.StatusBadRequest, err.Error())
	}
}

func editPage(page *models.Page) (*models.Page, error) {
	return_content := page.ReturnContent
	page.Path = GeneratePagePath(page.Title)
	repo := storage_repo.GetStorageRepo(context.Background())
	page, err := repo.EditPage(page.Id, page)
	if err != nil {
		return nil, err
	} else {
		if !return_content {
			page.Content = nil
		}
		page.CanEdit = true
		page.Url = GetPageUrl(page.Path)
		return page, nil
	}
}

func GetPageCount(access_token string) (int, error) {
	account, err := ValidateAccessToken(access_token)
	if err != nil {
		msg := fmt.Sprintf("Failed to validate access token: %s", err)
		fmt.Println(msg)
		return -1, errors.New(msg)
	}
	repo := storage_repo.GetStorageRepo(context.Background())
	page_list, err := repo.ListPages(account.Id, 1, 0)
	if err != nil {
		msg := fmt.Sprintf("Failed to list pages for account: %d", account.Id)
		fmt.Println(msg)
		return -1, errors.New(msg)
	}
	return page_list.TotalCount, nil
}

func GetPageList(c *gin.Context) {
	access_token := c.Query("access_token")
	t := c.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(t)
	if err != nil {
		msg := fmt.Sprintf("Wrong offset: %s", t)
		ReturnError(c, http.StatusBadRequest, msg)
		return
	}
	t = c.DefaultQuery("limit", "2")
	limit, err := strconv.Atoi(t)
	if err != nil {
		msg := fmt.Sprintf("Wrong limit: %s", t)
		ReturnError(c, http.StatusBadRequest, msg)
		return
	}

	account, err := ValidateAccessToken(access_token)
	if err != nil {
		msg := fmt.Sprintf("Failed to validate access token: %s", err)
		ReturnError(c, http.StatusForbidden, msg)
		return
	}
	repo := storage_repo.GetStorageRepo(context.Background())
	page_list, err := repo.ListPages(account.Id, limit, offset)
	if err != nil {
		msg := fmt.Sprintf("Failed to list pages for account: %d", account.Id)
		ReturnError(c, http.StatusBadRequest, msg)
		return
	}
	for i := 0; i < page_list.Count; i++ {
		page_list.Pages[i].CanEdit = true
		page_list.Pages[i].Url = GetPageUrl(page_list.Pages[i].Path)
	}
	ReturnSuccess(c, http.StatusOK, page_list)
}

func GetViews(c *gin.Context) {
	path := c.Param("path")
	var pv_req models.PageViewRequest
	err := c.ShouldBindQuery(&pv_req)
	if err != nil {
		msg := fmt.Sprintf("Failed to parse query: %s", err)
		ReturnError(c, http.StatusBadRequest, msg)
		return
	}

	repo := storage_repo.GetStorageRepo(context.Background())
	pv, err := repo.GetPageView(path, pv_req.Year, pv_req.Month, pv_req.Day, pv_req.Hour)
	if err == nil {
		res := gin.H{"views": pv}
		ReturnSuccess(c, http.StatusOK, res)
	} else {
		ReturnError(c, http.StatusBadRequest, err.Error())
	}
}
