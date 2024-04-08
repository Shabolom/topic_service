package tools

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"service_topic/internal/models"
)

func MultipartFormTopic(c *gin.Context, topic *models.Topic) (Claims, string, error) {
	var pathToFile string
	dataString := c.Request.FormValue("data")
	data := []byte(dataString)

	claims, err := ParseTokenClaims(c)

	err = json.Unmarshal(data, &topic)
	if err != nil {
		return Claims{}, "", err
	}

	for _, fileHeader := range c.Request.MultipartForm.File["file"] {

		file, err2 := fileHeader.Open()
		if err2 != nil {
			return Claims{}, "", err
		}

		path, err2 := MakeDir(file, "topic", fileHeader.Filename)
		if err2 != nil {
			return Claims{}, "", err2
		}
		pathToFile = path
	}

	return claims, pathToFile, nil
}
