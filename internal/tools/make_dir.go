package tools

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// MakeDirAndFile return path and error
func MakeDirAndFile(file multipart.File, dirName string, fileName string) (string, error) {
	err := os.MkdirAll(dirName+"_files", 0666)
	if err != nil {
		log.WithField("component", "rest").Debug(err)
		return "", err
	}

	unix := fmt.Sprintf("%v", time.Now().Unix())
	path := filepath.Join(dirName+"_files", unix+"."+fileName)

	targetFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.WithField("component", "tools").Debug(err)
		return "", err
	}
	defer targetFile.Close()

	_, err = io.Copy(targetFile, file)
	if err != nil {
		log.WithField("component", "tools").Debug(err)
		return "", err
	}

	return path, nil
}
