package main

import (
	"context"
	"errors"
	"strings"
)

func init() {
	prerequisites = [][]string{
		{"Git", "git"},
		{"JDK", "javac"},
		{"Android Studio", "studio64"},
		{"Flutter", "flutter"},
	}
}

func getGitWindowsLatestRelease(arch string) (string, string, error) {
	rel, _, err := ghClient.Repositories.GetLatestRelease(context.Background(), "git-for-windows", "git")
	panicIfErr(err)

	installerPrefix := "64-bit.exe"
	if strings.HasSuffix(arch, "86") {
		installerPrefix = "32-bit.exe"
	}

	for _, asset := range rel.Assets {
		if strings.HasSuffix(*asset.Name, installerPrefix) {
			return *asset.Name, *asset.BrowserDownloadURL, nil
		}
	}

	return "", "", errors.New("no git installer matched for this system")
}

func (i installer) downloadGit() {
	_, installerURL, err := getGitWindowsLatestRelease(i.arch)
	panicIfErr(err)
	downloadFile(installerURL)
}
