package commands

const (
	VERSION string = "--version"
	HELP string = "--help"

	VIDEO_1080 = "bestvideo[height<=1080]+bestaudio"
	VIDEO_720  = "bestvideo[height<=720]+bestaudio"
	VIDEO_480  = "bestvideo[height<=480]+bestaudio"

	EXTRACT_AUDIO = "-x"
	AUDIO_FORMAT = "--audio-format"
)