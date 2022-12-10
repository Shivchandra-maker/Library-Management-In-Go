package models

type Book struct {
	BookID        int    `json:"bookID"`
	AuthorID      int    `json:"authorID"`
	Auth          Author `json:"auth"`
	Title         string `json:"title"`
	Publication   string `json:"publication"`
	PublishedDate string `json:"publishedDate"`
}
