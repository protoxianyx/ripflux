package core

import (
	"ripflux/config/errors"
	"ripflux/core/models"
)

func BuildArgs(req models.DownloadRequest) []string {
	args := []string{}

	switch req.Format {
	case "video":
		args = append(args, "-f")

		switch req.Resolution {
		case "1080":
			args = append(args, "bestvideo[height<=1080]+bestaudio")
		case"720":
			args = append(args, "bestvideo[height<=720]+bestaudio")
		default:
			args = append(args, errors.MATCH_NOT_FOUND)
		}

	case "audio":
		args = append(args,
			"-x",
			"--audio-format",
			"mp3")
	}

	return args
}
