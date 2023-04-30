package schemas

import "github.com/gin-gonic/gin"

type BaseSchema interface {
	Validate(c *gin.Context, response *Response)
}

func schemaValidatorMacro(err error, response *Response) {
	if err != nil {
		response.Status, response.Message = 400, "Invalid Request Payload "+err.Error()
		panic("Invalid Request Body")
	}
}
