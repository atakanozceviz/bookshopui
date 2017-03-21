package model

type PageData struct {
	Title  string
	Active string
	User   string
	Type   string
	Books  Books
}

type Book struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Img       string `json:"img"`
	Price     string `json:"price"`
	WebSite   string `json:"website"`
}

type Books []Book
