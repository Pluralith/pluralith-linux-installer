package backend

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

type State struct {
	PluralithDir                string
	BinDir                      string
	DownloadProgress            float32
	DownloadProgressIncrementer chan float32
	DownloadComplete            bool
	InstallComplete             bool
}

// Instantiating State store
var StateStore = &State{}

func (S *State) InitState() {
	S.DownloadProgress = 0
	S.DownloadComplete = false
	S.InstallComplete = false
}

func (S *State) InitDirectories() error {
	functionName := "InitDirectories"

	// Get relevant directories
	home, homeErr := homedir.Dir()
	if homeErr != nil {
		return fmt.Errorf("failed to get home directory -> %v: %w", functionName, homeErr)
	}

	StateStore.PluralithDir = filepath.Join(home, "Pluralith")
	StateStore.BinDir = filepath.Join(StateStore.PluralithDir, "bin")

	// Create relevant directories
	dirErr := os.MkdirAll(StateStore.BinDir, 0700)
	if dirErr != nil {
		return fmt.Errorf("failed to create directories -> %v: %w", functionName, dirErr)
	}

	return nil
}
