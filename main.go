package main

import (
	"shorturl/internals"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	internals.SetupRoutes(router)

	router.Run()
}
