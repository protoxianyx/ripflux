package utils

import (
	"os"
	"path/filepath"
	"ripflux/config"
	"ripflux/utils/loggers"
)

func CleanTmp(tmpDirPath string) error {
	entries, err := os.ReadDir(tmpDirPath)
	if err != nil {
		return err
	}

	for _, entries := range entries {
		if entries.Name() == ".gitignore" {
			continue
		}

		targetPath := filepath.Join(tmpDirPath, entries.Name())

		if err := os.RemoveAll(targetPath); err != nil {
			loggers.MultiLogf(config.ERROR_LOG_FILE_PATH, true, "Failed to delete %s: %v", targetPath, err)
		}

		loggers.MultiLogf(config.COMBINED_LOG_FILE_PATH, true, "Deleted %s", targetPath)
	}

	return nil
}
