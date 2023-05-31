package api

type Tool struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	Editor      string `json:"editor"`
	Categorie   string `json:"categorie"`
}
type Tools []Tool

type Categorie struct {
	Name  string `json:"name"`
	Tools []Tool `json:"tools"`
}
type Categories []Categorie
