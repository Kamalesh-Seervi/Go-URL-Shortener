package main

import (
	"github.com/Kamalesh-Seervi/url/models"
	"github.com/Kamalesh-Seervi/url/server"
)

func main() {
	models.Setup()
	server.ServerListen()
}
