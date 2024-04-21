package Controllers

import (
	models "github.com/crudapi/Models"
	configs "github.com/crudapi/config"
	"github.com/gin-gonic/gin"
)

type clienteRequestBody struct {
	Name     string `json:"name"`
	Apellido string `json:"apellido"`
	Year     uint   `json:"year"`
}

func ClienteCreate(c *gin.Context) {

	body := clienteRequestBody{}

	c.BindJSON(&body)

	cliente := &models.Cliente{Name: body.Name, Apellido: body.Apellido, Year: body.Year}
	result := configs.DB.Create(&cliente)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, &cliente)
}

func ClienteGet(c *gin.Context) {
	var cliente []models.Cliente
	configs.DB.Find(&cliente)
	c.JSON(200, &cliente)

}

func ClienteGetById(c *gin.Context) {
	id := c.Param("id")
	var cliente models.Cliente
	configs.DB.First(&cliente, id)
	c.JSON(200, &cliente)

}

func ClienteUpdate(c *gin.Context) {

	id := c.Param("id")
	var cliente models.Cliente
	configs.DB.First(&cliente, id)

	body := clienteRequestBody{}
	c.BindJSON(&body)

	data := &models.Cliente{Name: body.Name, Apellido: body.Apellido, Year: body.Year}

	result := configs.DB.Model(&cliente).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": true, "message": "failed"})
		return
	}
	c.JSON(200, &cliente)
}

func ClienteDelete(c *gin.Context) {
	id := c.Param("id")
	var cliente models.Cliente
	configs.DB.Delete(&cliente, id)
	c.JSON(200, gin.H{"deleted": true})

}
