package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
)

func GetResponse[T Release | Repository](url string) (result []T) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		var message Error
		err = json.Unmarshal(body, &message)
		if err != nil {
			log.Fatal(err)
		}

		color.Red("%s", message.Message)
		os.Exit(1)
	}
	return
}

func DisplayReleases(username string, repository string) {
	url := fmt.Sprintf("%s/repos/%s/%s/releases?per_page=128", API, username, repository)
	releases := GetResponse[Release](url)

	if len(releases) == 0 {
		color.Red("No Releases")
		os.Exit(0)
	}

	var releasesDownloads int
	for i := range releases {
		release := releases[len(releases)-i-1]

		color.Magenta("Release Name: %s", release.Name)
		color.Green("Release Info")
		color.Cyan("Created at: %s", release.CreatedAt)
		color.Cyan("Published at: %s", release.PublishedAt)
		color.Cyan("Release Author: %s", release.Author.Name)

		if len(release.Assets) != 0 {
			color.Green("Download Info")

			var assetsDownloads int
			for _, asset := range release.Assets {
				assetsDownloads += asset.Downloads
				color.Cyan("%s (%d download tally)", asset.Name, asset.Downloads)
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
	url := fmt.Sprintf("%s/users/%s/repos", API, username)
	repositories := GetResponse[Repository](url)

	for _, repository := range repositories {
		color.Green("Repository Name: %s", repository.Name)
		color.Cyan("Pushed at: %s", repository.PushedAt)
		color.Cyan("Created at: %s", repository.CreatedAt)
		color.Cyan("Updated at: %s", repository.UpdatedAt)
	}

	color.Yellow("Repositories tally: %d", len(repositories))
}
