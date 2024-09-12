package main

import (
	"embed"
	"io/fs"
	"os"

	"github.com/BugzTheBunny/GoTH-Starter/routes"
	"github.com/gin-gonic/gin"
)

var (
	StaticAssetsFS embed.FS
)

func AssetsFS() fs.FS {
	if gin.Mode() == gin.DebugMode {
		return os.DirFS("static")
	} else {
		sub, err := fs.Sub(StaticAssetsFS, "static")
		if err != nil {
			panic(err)
		}
		return sub
	}
}

func main() {
	server := gin.Default()
	server.Static("/static", "static")
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
