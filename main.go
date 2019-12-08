package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v28/github"
)

func main() {
	client := github.NewClient(nil)
	LastReleases, _, _ := client.Repositories.ListReleases(context.Background(), "rabbitmq", "erlang-rpm", nil)
	for _, release := range LastReleases {
		assets, _, _ := client.Repositories.ListReleaseAssets(context.Background(), "rabbitmq", "erlang-rpm", release.GetID(), nil)
		for _, asset := range assets {
			fmt.Println(asset.GetName(), asset.GetBrowserDownloadURL())
		}
	}
}
