package main

import (
	"github.com/BugzTheBunny/GoTH-Starter/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
