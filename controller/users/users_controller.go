package users

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func GetUser(c *gin.Context) {
	log.Println("inside get")
	c.String(http.StatusNotImplemented, "implement me!")
}
