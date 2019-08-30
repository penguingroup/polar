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
	City     string `json:"city" form:"city" required:"true"`
	Category string `json:"category" form:"category" required:"true"`
	Page     int64  `json:"page" form:"page" required:"true"`
	Size     int64  `json:"size" form:"size" required:"true" default:"10"`
}
