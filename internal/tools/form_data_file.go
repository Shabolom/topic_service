package tools

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func FormDataFiles(c *gin.Context, dirName string) ([]string, error) {
	var pathToFile []string

	for _, fileHeader := range c.Request.MultipartForm.File["file"] {
		file, err := fileHeader.Open()
		if err != nil {
			log.WithField("component", "tools").Debug(err)
			return []string{}, err
		}

		path, err := MakeDirAndFile(file, dirName, fileHeader.Filename)
		if err != nil {
			return []string{}, err
		}
		pathToFile = append(pathToFile, path)
	}

	return pathToFile, nil
}
