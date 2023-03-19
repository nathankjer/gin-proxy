package controllers

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"crypto/md5"

	"github.com/gin-gonic/gin"
	"github.com/nathankjer/gin-proxy/db"
	"github.com/nathankjer/gin-proxy/models"
)

func Request(c *gin.Context) {
	cacheControl := c.Request.Header.Get("Cache-Control")
	body, _ := io.ReadAll(c.Request.Body)
	requestID := fmt.Sprintf("%x", md5.Sum([]byte(c.Request.Method+c.Request.Host+c.Request.URL.String()+string(body))))

	var request models.Request
	skipCache := cacheControl == "no-cache"

	if !skipCache {
		err := db.DB.Where("id = ?", requestID).First(&request).Error
		if err == nil {
			c.Data(request.ResponseStatus, "", request.ResponseBody)
			return
		}
	}

	// Perform request using proxy
	req, err := http.NewRequest(c.Request.Method, "http://localhost:3001"+c.Request.URL.Path+"?"+c.Request.URL.RawQuery, c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Save request
	request = models.Request{
		Id:             requestID,
		Method:         c.Request.Method,
		Host:           c.Request.Host,
		Path:           c.Request.URL.Path,
		ResponseStatus: res.StatusCode,
		ResponseBody:   responseBody,
		CreatedAt:      time.Now(),
	}
	if !skipCache {
		db.DB.Create(&request)
	}
	c.Data(res.StatusCode, "", responseBody)
}
