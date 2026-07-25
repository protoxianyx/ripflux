package adapters

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"ripflux/config"
	"ripflux/config/commands"
)

func GetVersion() {
	YTDLP := filepath.Join(config.BIN_DIR, config.YTDLP_EXE)

	cmd := exec.Command(YTDLP, commands.VERSION)
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
}