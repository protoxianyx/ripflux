package adapters

import (
	"fmt"
	"os"
	// "os"
	"os/exec"
	"path/filepath"

	// "ripflux/config"
	// "ripflux/config"
	"ripflux/config/commands"
	"ripflux/config/paths"
	"ripflux/utils/loggers"
	// "ripflux/core"
	// "ripflux/core/models"
)


var YTDLP string = filepath.Join(paths.BIN_DIR, paths.YTDLP_EXE)

func GetVersion() {
	
	cmd := exec.Command(YTDLP, commands.VERSION)
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	// fmt.Printf("The version of ytdlp: %s\n", string(output))
	loggers.Logf("The version of YTDLP: %s\n", string(output))
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