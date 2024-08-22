package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/savioxavier/termlink"
)

func GetResponse[T Release | Repository](url string, result *[]T) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		var message Error
		if err := json.Unmarshal(body, &message); err != nil {
			return err
		}
		return fmt.Errorf("%s", message.Message)
	}
	return nil
}

func DisplayReleases(username string, repository string) {
	var releases []Release
	url := fmt.Sprintf("%s/repos/%s/%s/releases?per_page=128", API, username, repository)
	err := GetResponse(url, &releases)
	if err != nil {
		color.Red("%s", err)
		os.Exit(1)
	}

	if len(releases) == 0 {
		color.Red("No Releases")
		os.Exit(0)
	}

	var releasesDownloads int
	for i := range releases {
		release := releases[len(releases)-i-1]

		color.Magenta("Release Name: %s", termlink.Link(release.Name, release.Url))
		color.Green("Release Info")
		color.Cyan("Release Author: %s", termlink.Link(release.Author.Name, release.Author.Url))
		color.Cyan("Created at: %s", release.CreatedAt)
		color.Cyan("Published at: %s", release.PublishedAt)

		if len(release.Assets) != 0 {
			color.Green("Download Info")

			var assetsDownloads int
			for _, asset := range release.Assets {
				assetsDownloads += asset.Downloads
				color.Cyan("%s (%d download tally)", termlink.Link(asset.Name, asset.Url), asset.Downloads)
			}

			releasesDownloads += assetsDownloads
			color.Green("Assets download tally: %d", assetsDownloads)
		}
	}

	if releasesDownloads != 0 {
		color.Yellow("Releases download tally: %d", releasesDownloads)
	}
}

func DisplayRepositories(username string) {
	var repositories []Repository
	url := fmt.Sprintf("%s/users/%s/repos", API, username)
	err := GetResponse(url, &repositories)
	if err != nil {
		color.Red("%s", err)
		os.Exit(1)
	}

	for _, repository := range repositories {
		color.Green("Repository Name: %s", termlink.Link(repository.Name, repository.Url))
		color.Cyan("Pushed at: %s", repository.PushedAt)
		color.Cyan("Created at: %s", repository.CreatedAt)
		color.Cyan("Updated at: %s", repository.UpdatedAt)
	}

	color.Yellow("Repositories tally: %d", len(repositories))
}
