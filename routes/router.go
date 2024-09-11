package routes

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func renderTemplate(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}

func RegisterRoutes(server *gin.Engine) {

	server.GET("/", homePage)

}
