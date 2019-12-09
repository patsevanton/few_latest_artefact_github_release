package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-ini/ini"
	"github.com/google/go-github/v28/github"
)

func main() {
	cfg, err := ini.Load("/etc/few_latest_artefact_github_release.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// Classic read of values, default section can be represented as empty string
	profile := cfg.Section("").Key("profile").String()
	repository := cfg.Section("").Key("repository").String()

	client := github.NewClient(nil)
	LastReleases, _, _ := client.Repositories.ListReleases(context.Background(), profile, repository, nil)
	for _, release := range LastReleases {
		assets, _, _ := client.Repositories.ListReleaseAssets(context.Background(), profile, repository, release.GetID(), nil)
		for _, asset := range assets {
			fmt.Println(asset.GetName(), asset.GetBrowserDownloadURL())
		}
	}
}
