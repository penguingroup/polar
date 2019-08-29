package models

type News struct {
	Id         int      `json:"id"`
	Title      string   `json:"title"`
	SubTitle   string   `json:"sub_title"`
	Weight     float32  `json:"weight"`
	Poster     string   `json:"poster"`
	Content    string   `json:"content"`
	Status     int      `json:"status"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
	Cities     []string `json:"cities"`
	Categories []string `json:"categories"`
}
