package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/Masterminds/semver/v3"
)

const githubAPI = "https://api.github.com/repos/arduino/arduino-cli/releases/latest"

type releaseInfo struct {
    TagName string `json:"tag_name"`
}

func getLatestVersion() (string, error) {
    resp, err := http.Get(githubAPI)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("GitHub API returned non-200 status: %d", resp.StatusCode)
    }

    var release releaseInfo
    if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
        return "", err
    }

    return release.TagName, nil
}

func main() {
    versionStr, err := getLatestVersion()
    if err != nil {
        log.Fatalf("Error fetching latest version: %v", err)
    }

    fmt.Printf("Reported version string: %s\n", versionStr)

    version, err := semver.NewVersion(versionStr)
    if err != nil {
        log.Fatalf("Error parsing version: %v", err)
    }

    fmt.Printf("Latest version of arduino-cli: %s\n", version)
}

