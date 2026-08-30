package config

import (
	"path/filepath"
	"ripflux/config/paths"
)

const (
	VIDEO = "Video"  
)


func logPathBuidler(logFileName string) string {

	logPath := filepath.Join(paths.CURRENT_LOCATION, paths.PATH_GO_BACK, paths.LOG_FOLDER, logFileName)

	return logPath
}

var OUTPUT_TEMPLATE_PATH string = filepath.Join(paths.TMP_FOLDER, "%(title)s.%(ext)s")
var OUTPUT_TMP_PATH string = filepath.Join(".", paths.PATH_GO_BACK, paths.TMP_FOLDER, "%(title)s.%(ext)s")
var YTDLP_BIN string = filepath.Join(paths.BIN_DIR, paths.ConstVars.YTDLP)

var COMBINED_LOG_FILE_PATH string = logPathBuidler(paths.LogFiles.COMBINED_LOGS)
var INPUT_LOG_FILE_PATH string = logPathBuidler(paths.LogFiles.INPUT_LOG)
var TEST_LOG_FILE_PATH string = logPathBuidler(paths.LogFiles.TEST_LOGS)
var ERROR_LOG_FILE_PATH string = logPathBuidler(paths.LogFiles.ERROR_LOGS)
var LOG_FILE_PATH string = logPathBuidler(paths.LogFiles.LOG)
var VERSION_INFO_FILEPATH string = logPathBuidler(paths.LogFiles.VERSION_LOG)


