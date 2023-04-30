package api

import (
	"App-Name/controllers"
	"github.com/gin-gonic/gin"
)

func CoreRoutes(engine *gin.Engine) {
	coreRouter := engine.Group("/")
	{
		coreRouter.POST("ping", controllers.HelloController)
		//coreRouter.POST("login", jwtMiddleware.LoginHandler)
	}
}
