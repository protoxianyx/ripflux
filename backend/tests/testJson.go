package json_tests

import (
	"encoding/json"

	"log"

	// "fmt"
	"os"
	"ripflux/config"
	"ripflux/core/models"
)

func Json_test() models.DownloadRequest {

	data, err := os.Open(config.TEST_DATA)
	if err != nil {
		wd, _ := os.Getwd()
		log.Printf("Error: %s || Path: %s", err, wd)
	}
	defer data.Close()

	var req models.DownloadRequest

	err = json.NewDecoder(data).Decode(&req)

	return req

}
