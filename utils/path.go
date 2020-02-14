package utils

import (
	"fmt"
	"os"
	"runtime"
)

func UserHomeDir() string {
	fmt.Println()
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func GetStoragePath() string {
	if runtime.GOOS == "windows" {
		appDataPath := os.Getenv("APPDATA")
		if appDataPath == "" {
			appDataPath = fmt.Sprintf("%s\\AppData\\Roaming", UserHomeDir())
		}
		return fmt.Sprintf("%s\\QS", appDataPath)
	} else if runtime.GOOS == "linux" {
		return fmt.Sprintf("%s\\.config\\QS", UserHomeDir()) // LINUX
	} else if runtime.GOOS == "darwin" {
		return "~\\Library\\Application Support\\QS"
	}
	return "./QS" // Default
}
