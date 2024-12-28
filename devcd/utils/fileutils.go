package utils

import (
	"devcd/logger"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func CreateTmpDirWithTS(tmpDir, tmpFolder string) error {

	dateStamp := time.Now().Format("20060102")
	timeStamp := time.Now().Format("150405")
	tempDir := fmt.Sprintf("%s/%s_%s_%s", tmpDir, tmpFolder, dateStamp, timeStamp)
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		err := os.MkdirAll(tempDir, os.ModePerm)
		if err != nil {
			logger.Error("Error creating temporary directory", "tempDir", tempDir, "error", err)
			return err
		} else {
			logger.Info("Created temporary directory", "tempDir", tempDir)
		}
	} else {
		logger.Warn("Temporary directory already exists", "tempDir", tempDir)
	}
	return nil
}

func CreateTmpDir(tmpDir string) error {

	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		err := os.MkdirAll(tmpDir, os.ModePerm)
		if err != nil {
			logger.Error("Error creating temporary directory", "tempDir", tmpDir, "error", err)
			return err
		} else {
			logger.Info("Created temporary directory", "tempDir", tmpDir)
		}
	} else {
		logger.Warn("Temporary directory already exists", "tempDir", tmpDir)
	}
	return nil
}

// copyFile copies the source file to the destination.
func CopyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func HandleDestination(src, dest string) error {
	fileInfo, err := os.Stat(dest)
	if err != nil {
		if os.IsNotExist(err) {
			// Destination doesn't exist, proceed with copy.
			return CopyFile(src, dest)
		}
		return err
	}

	if fileInfo.IsDir() {
		// Destination is a directory, remove it and then copy.
		if err := os.RemoveAll(dest); err != nil {
			return err
		}
		return CopyFile(src, dest)
	}

	// Destination is an existing file, skip the copy.
	logger.Info("Destination file already exists. Skipping copy.")
	return nil
}

func CopyDirectory(srcDir, destDir string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		return CopyFile(path, destPath)
	})
}

func RemoveDirectory(destDir string) error {
	files, err := os.ReadDir(destDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		err = os.RemoveAll(filepath.Join(destDir, file.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveDirectoryExcludeGitIgnore(destDir string) error {
	files, err := os.ReadDir(destDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".gitkeep") {
			err = os.RemoveAll(filepath.Join(destDir, file.Name()))
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func GetCwd() (string, error) {

	cwdPath, err := os.Executable()
	if err != nil {
		logger.Error("Error getting CWD")
		return "", err
	}

	cwd := filepath.Dir(cwdPath)
	return cwd, nil
}
