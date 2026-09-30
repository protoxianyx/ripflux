package models

type DownloadRequestModel struct {
	URL        string `json:"url"`
	Format     string `json:"format"`
	Resolution string `json:"resolution"`
	DryRun     string `json:"dry_run,omitempty"`
}

type CommandStructModel struct {
	Command string
	CmdArgs []string
}

type SubCommandModel struct {
	Flag     string
	FlagArgs []string
}

func (*SubCommandModel) buildCmdArgs(command string, args SubCommandModel) CommandStructModel {

	return CommandStructModel{command, args.FlagArgs}
}
