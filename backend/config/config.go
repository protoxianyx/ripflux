package config

import (
	"path/filepath"
	"ripflux/config/paths"
)

const (
	TEMP_OUTPUT string = "temp_output"
)

var OUTPUT_TEMPLATE_PATH string = filepath.Join(paths.TMP_FOLDER, "%(title)s.%(ext)s")
var OUTPUT_TMP_PATH string = filepath.Join(".", paths.PATH_GO_BACK, paths.TMP_FOLDER, "%(title)s.%(ext)s")
var BACKEND_LOG_FILE_PATH string = filepath.Join(paths.PATH_GO_BACK, paths.LOG_FOLDER, paths.BACKEND_LOG_FILE)
