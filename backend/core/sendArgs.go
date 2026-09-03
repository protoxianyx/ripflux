package core

import "ripflux/core/adapters"

func Downloader(args []string) error {

	return adapters.Download(args)
}
