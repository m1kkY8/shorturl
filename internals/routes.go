package internals

import (
	"net/http"
	"shorturl/views"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.StaticFile("/style.css", "./public/style.css")

	router.NoRoute(func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/404")
	})

	router.GET("/404", func(ctx *gin.Context) {
		render(ctx, 200, views.NotFound404())
	})

	router.GET("/", func(ctx *gin.Context) {
		render(ctx, 200, views.Index())
	})

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "https://www.google.com")
	})
}
