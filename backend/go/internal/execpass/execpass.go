package execpass

import (
	"os"
	"os/exec"
)

func ExecPass(url string) string {
	cmd := exec.Command("../bin/yt-dlp.exe", url)

	cmd.Dir = "../../download"

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	return url + "passed"
}
