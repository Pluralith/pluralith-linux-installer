package backend

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"
)

func LaunchPluralith() error {
	functionName := "LaunchPluralith"

	time.Sleep(time.Millisecond * 1000)
	appImagePath := filepath.Join(StateStore.BinDir, "Pluralith.AppImage")

	// Open Pluralith
	cmd := exec.Command(appImagePath)
	if err := cmd.Start(); err != nil {
		time.Sleep(200 * time.Millisecond)
		return fmt.Errorf("failed to launch Pluralith -> %v: %w", functionName, err)
	}

	time.Sleep(1000 * time.Millisecond)
	return nil
}
