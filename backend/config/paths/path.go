package paths

const (
	BIN_DIR    string = "../bin"
	FFMPEG_EXE string = " "

	TEST_DATA string = "tests/data.json"

	LOG_FOLDER       string = "logs"
	BACKEND_LOG_FILE string = "backend.log"

	TMP_FOLDER       string = "tmp"
	PATH_GO_BACK     string = ".."
	CURRENT_LOCATION string = "."

	// YTDLPPath string = "../bin/yt-dlp.exe"
)

type constVars struct {
	YTDLP string
}

var ConstVars = constVars{
	// TEMP_OUTPUT: "temp_output",
	YTDLP: "yt-dlp.exe",
	// INSTALLER_LOG_FILE string = "installer.log"
}
