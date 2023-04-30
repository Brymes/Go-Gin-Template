package schemas

import (
	"App-Name/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Response struct {
	Status  int         `json:"-"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (response *Response) HandlePanic(err interface{}, logger *log.Logger) {
	if utils.HandlePanicMacro(err, logger) {
		if response.Status == 0 {
			response.Status = 500
		}
		response.Success = false
		if response.Message == "" {
			response.Message = "Error processing your request at the current time. Kindly contact support"
		}
	}
}

func (response *Response) SendResponse(logger *log.Logger, c *gin.Context) {
	err := recover()
	response.HandlePanic(err, logger)

	if !response.Success || response.Status != 0 {
		c.IndentedJSON(response.Status, response)
	} else {
		response.Success = true
		c.IndentedJSON(http.StatusOK, response)
	}

}

func (response *Response) ValidatorMacro(c *gin.Context, validator BaseSchema) {
	validator.Validate(c, response)
}
func (response *Response) HandleError(err error, code int) {
	if err != nil {
		response.Status = code
	}
}

func (response *Response) HandleDbErrorMacro(result *gorm.DB, logger *log.Logger, message string) {
	response.Message, response.Status = utils.HandleDbError(result, logger)

	if response.Status != 0 {
		response.Success, response.Data = false, nil
	} else {
		response.Message, response.Success = message, true
	}
}
