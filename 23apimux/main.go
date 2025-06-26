package main

// Model for Course and Author - course.go and author.go
type Course struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Duration string  `json:"duration"`
	Price    float64 `json:"price"`
	Author   *Author `json:"author"`
}

type Author struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

// Fake DB
var courses = []Course{}

// Middleware for empty fields
func (c *Course) IsEmpty() bool {
	return c.ID == "" && c.Name == ""
}

func main() {

}
