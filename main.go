package main

import (
	"net/http"

	routes "github.com/crudapi/Routes"
	configs "github.com/crudapi/config"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectToDB()
}

// conexion al gin server
func main() {

	r := gin.Default()

	routes.ClienteRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from server.",
		})
	})
	r.Run()
}
