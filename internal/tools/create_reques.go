package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRequest(c *gin.Context, method string, path string) (*http.Response, error) {
	request, err := http.NewRequest(method, path, c.Request.Body)
	if err != nil {
		return &http.Response{}, err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return &http.Response{}, err
	}

	c.Writer.Header().Set("Authorization", response.Header.Get("Authorization"))

	return response, nil
}
