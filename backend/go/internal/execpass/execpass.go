package execpass

import (
	
	"os"
	"os/exec"
)

func ExecPass(url string) string {
	cmd := exec.Command("../../bin/yt-dlp.exe", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		panic(err)	
	}

	return url + "passed"
}