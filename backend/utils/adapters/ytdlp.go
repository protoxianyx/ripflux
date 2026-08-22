package adapters

import (
	"fmt"
	"os"
	"strings"

	// "os"
	"os/exec"
	"path/filepath"

	// "ripflux/config"
	// "ripflux/config"
	"ripflux/config"
	"ripflux/config/commands"
	"ripflux/config/paths"
	"ripflux/utils/loggers"
	// "ripflux/core"
	// "ripflux/core/models"
)


var YTDLP string = filepath.Join(paths.BIN_DIR, config.YTDLP_EXE)

func GetVersion() (string, error) {
	
	cmd := exec.Command(YTDLP, commands.VERSION)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("could not get yt-dlp version: %w", err)
	}

	version := strings.TrimSpace(string(output))
	// fmt.Printf("The version of ytdlp: %s\n", string(output))
	loggers.Logf("The version of YTDLP: %s\n", version)

	return version, nil

}

func Run() {
	
	fmt.Println()
}

func Download(args []string) {
	cmd := exec.Command(YTDLP, args...)
	

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		loggers.Logf("Execution failed: %v\n==========END==========\n", err)
	}

	loggers.Log("Download completed successfully!\n==========END==========\n")

}