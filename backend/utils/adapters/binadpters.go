package adapters

import (
	"os/exec"
	"ripflux/models"
)

type ytdlp_io struct {
}

func ExecuteYTDLP(args models.CommandStructModel) error {
	cmd := exec.Command(args.Command, args.CmdArgs...)

	cmd.Run()

	return nil
}
