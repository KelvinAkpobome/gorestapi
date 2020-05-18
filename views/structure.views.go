package views

// CreatePost is the structure of the database.
type CreatePost struct {
	ID int			`json:"id"`
	Author string		`json:"author"`
	Name string 	`json:"name"`
	PublishedAt string `json:"published_at"`
}

// UpdatePost for PUT request to the database.
type UpdatePost struct {
	Name string 	`json:"name"`
}
