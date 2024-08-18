package lib

const API = "https://api.github.com"

type Error struct {
	Message string `json:"message"`
}

type Author struct {
	Name string `json:"login"`
}

type Asset struct {
	Name      string `json:"name"`
	Downloads int    `json:"download_count"`
}

type Release struct {
	Name        string  `json:"name"`
	Author      Author  `json:"author"`
	Assets      []Asset `json:"assets"`
	CreatedAt   string  `json:"created_at"`
	PublishedAt string  `json:"published_at"`
}

type Repository struct {
	Name      string `json:"name"`
	PushedAt  string `json:"pushed_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
