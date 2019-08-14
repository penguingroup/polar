package main

import (
	"github.com/karldoenitz/Tigo/TigoWeb"
	"polar/handlers"
)

var urls = []TigoWeb.Router{
	{"/ping", &handlers.PingHandler{}, nil},
}

func main() {
	application := TigoWeb.Application{IPAddress: "0.0.0.0", Port: 8080, UrlRouters: urls}
	application.Run()
}
