package main

import (
	"awesomeProject/book/v3/config"
	"awesomeProject/book/v3/exception"
	"awesomeProject/book/v3/handlers"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "application.yaml"
	}
	config.LoadConfigfromYaml(path)

	server := gin.New()
	server.Use(gin.Logger(), exception.Recovery())

	handlers.Book.Registry(server)

	ac := config.C().Application

	if err := server.Run(fmt.Sprintf("%s:%d", ac.Host, ac.Port)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
