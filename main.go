package main

import (
	"ReshmaKolekar/bookstore_users/app"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	app.StartApplication()
}
