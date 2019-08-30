package main

import (
	"github.com/karldoenitz/Tigo/TigoWeb"
	"polar/global/config"
	"polar/handlers"
)

var urls = []TigoWeb.Router{
	{"/ping", &handlers.PingHandler{}, nil},
	{"/api/header/category", &handlers.CategoryHandler{}, nil},
	{"/api/header/city", &handlers.CityHandler{}, nil},
	{"/api/body/news", &handlers.NewsHandler{}, nil},
}

func main() {
	application := TigoWeb.Application{ConfigPath: config.GetServerConfig(), UrlRouters: urls}
	application.Run()
}
