package main

import (
	"github.com/aminkamal/golang_test/internal/service"
)

func main() {
	svc := service.New()
	svc.AddRoutes()
	svc.Router.Run()
}
