package routes

import (
	"github.com/crudapi/Controllers"
	"github.com/gin-gonic/gin"
)

func ClienteRouter(router *gin.Engine) {

	routes := router.Group("api/v1/clients")
	routes.POST("", Controllers.ClienteCreate)
	routes.GET("", Controllers.ClienteGet)
	routes.GET("/:id", Controllers.ClienteGetById)
	routes.PUT("/:id", Controllers.ClienteUpdate)
	routes.DELETE("/:id", Controllers.ClienteDelete)

}
