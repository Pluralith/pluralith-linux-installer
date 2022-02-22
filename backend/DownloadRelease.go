package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func getReleaseUrl() (string, error) {
	functionName := "getReleaseUrl"

	request, _ := http.NewRequest("GET", "https://api.pluralith.com/v1/dist/download/ui", nil)

	queryString := request.URL.Query()
	queryString.Add("os", "linux")
	request.URL.RawQuery = queryString.Encode()

	// Execute get version request
	client := &http.Client{}
	response, responseErr := client.Do(request)

	if responseErr != nil || response.StatusCode != 200 {
		return "", fmt.Errorf("fetching download link failed -> %v: %w", functionName, responseErr)
	}

	// Parse request body
	var bodyObject map[string]interface{}
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	parseErr := json.Unmarshal(bodyBytes, &bodyObject)
	if parseErr != nil {
		return "", fmt.Errorf("parsing request result failed -> %v: %w", functionName, responseErr)
	}

	bodyData := bodyObject["data"].(map[string]interface{})
	return bodyData["url"].(string), nil
}

func updateDownloadProgress(binPath string, totalSize float32) {
	for {
		newFile, err := os.Open(binPath)
		if err != nil {
			log.Fatal(err)
		}

		// Get current download file size
		currentInfo, err := newFile.Stat()
		if err != nil {
			log.Fatal(err)
		}
		currentSize := currentInfo.Size()
		if currentSize == 0 {
			currentSize = 1
		}

		// Update download progess for progress bar
		StateStore.DownloadProgressIncrementer <- float32(currentSize)/float32(totalSize) - StateStore.DownloadProgress

		// Break download loop if download complete
		if StateStore.DownloadProgress >= 1 {
			break
		}

		// Debounce progress bar update
		time.Sleep(time.Millisecond * 100)
	}
}

func DownloadRelease() error {
	functionName := "DownloadRelease"

	// Get release URL from API
	url, urlErr := getReleaseUrl()
	if urlErr != nil {
		return urlErr
	}

	// Construct AppImage path
	appImagePath := filepath.Join(StateStore.BinDir, "Pluralith.AppImage")

	// Create bin file
	newFile, createErr := os.Create(appImagePath)
	if createErr != nil {
		newFile.Close()
		return fmt.Errorf("failed to create binary on file system -> %v: %w", functionName, createErr)
	}

	defer newFile.Close()

	// Get latest version
	response, getErr := http.Get(url)
	if getErr != nil {
		response.Body.Close()
		return fmt.Errorf("downloading latest AppImage failed -> %v: %w", functionName, getErr)
	}

	defer response.Body.Close()

	// Get total size from http header
	totalSize, sizeErr := strconv.Atoi(response.Header.Get("Content-Length"))
	if sizeErr != nil {
		return fmt.Errorf("getting file size failed -> %v: %w", functionName, sizeErr)
	}

	// Update download progress for UI progress bar
	go updateDownloadProgress(appImagePath, float32(totalSize))

	// Write to file and progress bar
	if _, writeErr := io.Copy(io.MultiWriter(newFile), response.Body); writeErr != nil {
		return fmt.Errorf("downloading latest AppImage failed -> %v: %w", functionName, getErr)
	}

	// Make AppImage executable
	chmodErr := os.Chmod(appImagePath, 0700)
	if chmodErr != nil {
		return fmt.Errorf("making binary executable failed -> %v: %w", functionName, getErr)
	}

	return nil
}
