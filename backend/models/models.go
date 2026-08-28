package models

type DownloadRequestModel struct {
	URL        string `json:"url"`
	Format     string `json:"format"`
	Resolution string `json:"resolution"`
}

type ExecuteCommandModel struct {
	Command string
	CmdArgs []string

	Flag    string
	FlagArg string
}

type ServiveIOModel struct {
	Input ExecuteCommandModel
	Output string
    Logs string
}
