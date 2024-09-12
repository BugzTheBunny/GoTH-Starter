package routes

import (
	"net/http"

	"github.com/BugzTheBunny/GoTH-Starter/views/layout"
	"github.com/gin-gonic/gin"
)

func homePage(ctx *gin.Context) {
	renderTemplate(ctx, http.StatusOK, layout.CreateIndex("Bugz The Bunny"))
}
