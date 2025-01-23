package models

type Post struct {
	ID                 string `json:"id"`
	Title              string `json:"title"`
	Content            string `json:"content"`
	Date               string `json:"date"`
	Category           string `json:"category"`
	MetaTagTitle       string `json:"meta_tag_title"`
	MetaTagDescription string `json:"meta_tag_description"`
	PostImage          string `json:"post_image"`
	PostBackground     string `json:"post_background"`
	Author             string `json:"author"`
	Keywords           string `json:"keywords"`
}
