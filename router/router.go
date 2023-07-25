package router

import "github.com/gin-gonic/gin"

func Initialize() {

	Router := gin.Default()

	initializeRoutes(Router)

	Router.Run(":8080")
}