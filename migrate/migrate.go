package main

import (
	"github.com/AkshayAwate/go-crud/initializers"
	"github.com/AkshayAwate/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}
func main() {
	initializers.DB.AutoMigrate(&models.Post{})

}
