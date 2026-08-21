package json_tests

import (
	"encoding/json"

	"log"

	// "fmt"
	"os"
	// "ripflux/config"
	"ripflux/config/paths"
	"ripflux/core/models"
)

func Json_test() models.DownloadRequestModel {

	data, err := os.Open(paths.TEST_DATA)
	if err != nil {
		wd, _ := os.Getwd()
		log.Printf("Error: %s || Path: %s", err, wd)
	}
	defer data.Close()

	var req models.DownloadRequestModel

	err = json.NewDecoder(data).Decode(&req)

	return req

}
