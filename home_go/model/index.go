package model

type Body struct {
	Heading    string   `json:"heading"`
	Paragraphs []string `json:"paragraphs"`
}

type Link struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type IndexPage struct {
	Title   string `json:"title"`
	Heading string `json:"heading"`
	Main    Body   `json:"main"`
	Links   []Link `json:"links"`
}

