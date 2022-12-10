package models

type Author struct {
	AuthID    int    `json:"authID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Dob       string `json:"dob"`
	PenName   string `json:"penName"`
}
