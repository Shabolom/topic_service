package tools

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
)

func MultipartForm(c *gin.Context, obj interface{}, dirName string) (Claims, string, error) {
	var pathToFile []string
	dataString := c.Request.FormValue("data")
	data := []byte(dataString)

	claims, err := ParseTokenClaims(c)

	err = json.Unmarshal(data, &obj)

	if err != nil {
		fmt.Println(dataString)
		log.WithField("component", "tools").Debug(err)
		return Claims{}, "", err
	}

	for _, fileHeader := range c.Request.MultipartForm.File["file"] {

		file, err2 := fileHeader.Open()
		if err2 != nil {
			log.WithField("component", "tools").Debug(err2)
			return Claims{}, "", err
		}

		path, err2 := MakeDirAndFile(file, dirName, fileHeader.Filename)
		if err2 != nil {
			return Claims{}, "", err2
		}
		pathToFile = append(pathToFile, path)
	}

	path := strings.Join(pathToFile, "(space)")

	return claims, path, nil
}
