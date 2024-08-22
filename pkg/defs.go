package pkg

const API = "https://api.github.com"

type Error struct {
	Message string `json:"message"`
}

type Author struct {
	Url  string `json:"html_url"`
	Name string `json:"login"`
}

type Asset struct {
	Url       string `json:"browser_download_url"`
	Name      string `json:"name"`
	Downloads int    `json:"download_count"`
}

type Release struct {
	Url         string  `json:"html_url"`
	Name        string  `json:"name"`
	Author      Author  `json:"author"`
	Assets      []Asset `json:"assets"`
	CreatedAt   string  `json:"created_at"`
	PublishedAt string  `json:"published_at"`
}

type Repository struct {
	Url       string `json:"html_url"`
	Name      string `json:"name"`
	PushedAt  string `json:"pushed_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
