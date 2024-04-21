package main

import (
	models "github.com/crudapi/Models"
	configs "github.com/crudapi/config"
)

func init() {
	configs.ConnectToDB()
}

func main() {

	configs.DB.AutoMigrate(&models.Cliente{})
}
