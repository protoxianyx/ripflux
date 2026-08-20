package adapters

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"ripflux/config"
	"ripflux/config/commands"
	// "ripflux/core"
	// "ripflux/core/models"
)


var YTDLP string = filepath.Join(config.BIN_DIR, config.YTDLP_EXE)

func GetVersion() {
	
	cmd := exec.Command(YTDLP, commands.VERSION)
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Printf("The version of ytdlp: %s\n", string(output))
}

func Run() {
	
	fmt.Println()
}

func Download(args []string) {
	// cmd := exec.Command(YTDLP, )

	fmt.Println(args)
}