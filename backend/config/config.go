package config

import (
	"path/filepath"
	"ripflux/config/paths"
)

const (
	TEMP_OUTPUT        string = "temp_output"
	YTDLP_EXE          string = "yt-dlp.exe"
	INSTALLER_LOG_FILE string = "installer.log"
)

func logPathBuidler(logFileName string) string {

	logPath := filepath.Join(paths.CURRENT_LOCATION, paths.PATH_GO_BACK, paths.LOG_FOLDER, logFileName)

	return logPath
}

var OUTPUT_TEMPLATE_PATH string = filepath.Join(paths.TMP_FOLDER, "%(title)s.%(ext)s")
var OUTPUT_TMP_PATH string = filepath.Join(".", paths.PATH_GO_BACK, paths.TMP_FOLDER, "%(title)s.%(ext)s")

var COMBINED_LOG_FILE_PATH string = logPathBuidler(paths.COMBINED_LOGS)
var INPUT_LOG_FILE_PATH string = logPathBuidler(paths.INPUT_LOG)
