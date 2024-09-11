package routes

import (
	"net/http"

	"github.com/BugzTheBunny/GoTH-Starter/views"
	"github.com/gin-gonic/gin"
)

func homePage(ctx *gin.Context) {
	renderTemplate(ctx, http.StatusOK, views.CreateIndex("Bugz The Bunny"))
}
