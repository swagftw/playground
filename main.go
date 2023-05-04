package main

import (
	"errors"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	filePath := "/Users/swag/test/asd.txt"
	// get dir from filepath
	splits := strings.Split(filePath, "/")

	// remove filename
	splits = splits[:len(splits)-1]

	// join dir
	dirPath := strings.Join(splits, "/")

	// create directory if it doesn't exist
	if _, err := os.Stat(dirPath); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(dirPath, 0700)
		if err != nil {
			log.WithError(err).Error("failed to create directory")

			return
		}
	}

	err := os.WriteFile(filePath, []byte("hello world"), 0644)
	if err != nil {
		log.Errorf("error writing file %s to disk: %v", filePath, err)

		return
	}

	log.Debugf("Successfully wrote payload to file %s", filePath)

	return
}
