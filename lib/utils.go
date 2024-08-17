package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
)

func GetReleases(username string, repository string) (releases []Release) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases?per_page=128", username, repository)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &releases)
	if err != nil {
		var message Error
		err = json.Unmarshal(body, &message)
		if err != nil {
			log.Fatal(err)
		}

		color.Red("%s", message.Message)
	}
	return
}

func ShowReleases(username string, repository string) {
	releases := GetReleases(username, repository)

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
