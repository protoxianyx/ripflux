package adapters

import (
	"fmt"
	"os/exec"
	"ripflux/models"
)

type ytdlp_output struct {
	IsSucessful bool
	Error       error
}

func ExecuteYTDLP(args models.CommandStructModel) ytdlp_output {
	cmd := exec.Command(args.Command, args.CmdArgs...)
	// output, err := cmd.CombinedOutput()

	if err := cmd.Run(); err != nil {
		return ytdlp_output{false, fmt.Errorf("Something went wrong during executing")}
	}

	return ytdlp_output{true, nil}
}
