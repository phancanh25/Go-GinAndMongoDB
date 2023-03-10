package main

import (
	"GinAndMongoDB/configs"
	"GinAndMongoDB/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()
	routes.UserRoute(router)

	router.Run("localhost:6000")
}
