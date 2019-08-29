package global

type Category struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type City struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type NewsRequest struct {
	City     string `json:"city" required:"true"`
	Category string `json:"category" required:"true"`
	Page     int64  `json:"page" required:"true"`
	Size     int64  `json:"size" required:"true" default:"10"`
}
