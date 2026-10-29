package commands

type presets struct {
	VERSION string
	HELP    string

	VIDEO_1080 string
	VIDEO_720  string
	VIDEO_480  string
	VIDEO_MAX  string
	MP3        string
	WAV        string
}

type flags struct {
	VERSION string
	HELP    string

	AUDIO_EXTRACT string
	AUDIO_FORMAT  string

	FORMAT      string
	FORMAT_SORT string
	OUTPUT      string
}


type inputOptions struct {
	VIDEO string
	AUDIO string
	RESOLUTION int
}

var YTDLP_PRESETS = presets{

	VIDEO_1080: "bestvideo[height<=1080]+bestaudio",
	VIDEO_720:  "bestvideo[height<=720]+bestaudio",
	VIDEO_480:  "bestvideo[height<=480]+bestaudio",
	VIDEO_MAX:  "bestvideo+bestaudio/best",

	MP3: "mp3",
	WAV: "wav",
}

var YTDLP_FLAGS = flags{
	VERSION: "--version",
	HELP:    "--help",

	AUDIO_EXTRACT: "-x",
	AUDIO_FORMAT:  "--audio-format",

	FORMAT:      "--format",
	FORMAT_SORT: "--format-sort",

	OUTPUT: "--output",
}

var IMPUT_OPTIONS = inputOptions{
	VIDEO: "video",
	AUDIO: "audio",
}
