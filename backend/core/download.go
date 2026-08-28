package core

import (
	"ripflux/config"
	"ripflux/config/commands"
	"ripflux/config/errors"
	"ripflux/models"
	"ripflux/utils/adapters"
)

func BuildDownloadArgs(req models.DownloadRequestModel) []string {
	args := []string{}

	args = append(args, req.URL)
	switch req.Format {
	case "Video":
		args = append(args, commands.YTDLP_FLAGS.FORMAT)

		switch req.Resolution {
		case "1080":
			args = append(args, "bestvideo[height<=1080]+bestaudio")
		case "720":
			args = append(args, "bestvideo[height<=720]+bestaudio")
		default:
			args = append(args, errors.UserSide.MATCH_NOT_FOUND)
		}

	case "Audio":
		args = append(args,
			commands.YTDLP_FLAGS.AUDIO_EXTRACT,
			commands.YTDLP_FLAGS.AUDIO_EXTRACT,
			commands.YTDLP_PRESETS.MP3)
	}

	args = append(args, commands.YTDLP_FLAGS.OUTPUT, config.OUTPUT_TMP_PATH)

	return args
}

func Downloader(args []string) {
	adapters.Download(args)

}
