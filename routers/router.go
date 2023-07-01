package routers

import (
	"blogweb_gin/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	// register
	router.GET("/register", controllers.RegisterGet)
	return router
}
