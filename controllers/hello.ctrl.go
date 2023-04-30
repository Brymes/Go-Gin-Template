package controllers

import (
	"App-Name/config"
	"App-Name/models"
	"App-Name/schemas"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
)

func HelloController(c *gin.Context) {
	var (
		request   models.HelloModel
		validator schemas.HelloValidator
	)

	reqBuffer, reqLogger := config.InitRequestLogger("Create User")
	defer log.Println(reqBuffer)

	response := schemas.Response{Success: true}
	defer response.SendResponse(reqLogger, c)

	validator.Validate(c, &response)

	//Call Controller
	_ = c.ShouldBindBodyWith(&request, binding.JSON)
	response.Message, response.Success = request.SayHello(), true
}
