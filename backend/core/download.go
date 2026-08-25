package core

import (
	"ripflux/config"
	"ripflux/config/errors"
	"ripflux/core/models"
	"ripflux/utils/adapters"
)

func BuildArgs(req models.DownloadRequestModel) []string {
	args := []string{}

	// new args dsad

	args = append(args, req.URL)
	switch req.Format {
	case "Video":
		args = append(args, "-f")

		switch req.Resolution {
		case "1080":
			args = append(args, "bestvideo[height<=1080]+bestaudio")
		case "720":
			args = append(args, "bestvideo[height<=720]+bestaudio")
		default:
			args = append(args, errors.MATCH_NOT_FOUND)
		}

	case "Audio":
		args = append(args,
			"-x",
			"--audio-format",
			"mp3")
	}

	args = append(args, "--output", config.OUTPUT_TMP_PATH)

	return args
}

func Downloader(args []string) {
	adapters.Download(args)

}
