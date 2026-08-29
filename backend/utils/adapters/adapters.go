package adapters

import (
	"fmt"
	"os"
	"strings"

	"os/exec"
	"path/filepath"

	"ripflux/config"
	"ripflux/config/commands"
	"ripflux/config/paths"
	"ripflux/utils/loggers"
)

var YTDLP string = filepath.Join(paths.BIN_DIR, config.YTDLP_EXE)

func GetVersion() (string, error) {
	_, err := os.Stat(YTDLP)
	if err != nil {
		if os.IsNotExist(err) {
			return "", loggers.MTaskLog(config.ERROR_LOG_FILE_PATH, true, "YTDLP is missing")
		}

		return "", loggers.MTaskLog(config.ERROR_LOG_FILE_PATH, true, "Could not access YTDLP")
	}

	cmd := exec.Command(YTDLP, commands.YTDLP_FLAGS.VERSION)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("could not get yt-dlp version: %w", err)
	}

	version := strings.TrimSpace(string(output))
	loggers.MultiLogf(config.VERSION_INFO_FILEPATH, true, "This is the neew version line: %s\n", version)
	loggers.Logf("The version of YTDLP: %s\n", version)

	return version, nil

}

func Run() {

	fmt.Println()
}

func Download(args []string) error {
	cmd := exec.Command(YTDLP, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		loggers.Logf("Execution failed: %v\n==========END==========\n", err)
		// err := fmt.Sprintf("ytdlp execution failed: %v", err.Error())
		return loggers.MTaskLog(config.ERROR_LOG_FILE_PATH, true, err.Error())
	}

	loggers.Log("Download completed successfully!\n==========END==========\n")

	return nil
}

func UpdateInternally(args []string) error {
	cmd := exec.Command(YTDLP, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		loggers.Logf("Execution failed: %v\n==========END==========\n", err)
		// err := fmt.Sprintf("ytdlp execution failed: %v", err.Error())
		return loggers.MTaskLog(config.ERROR_LOG_FILE_PATH, true, err.Error())
	}

	loggers.Log("Download completed successfully!\n==========END==========\n")

	return nil
}
