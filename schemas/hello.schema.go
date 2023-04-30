package schemas

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type HelloValidator struct {
	Name string `json:"name" binding:"required"`
}

func (h *HelloValidator) Validate(c *gin.Context, response *Response) {
	err := c.ShouldBindBodyWith(&h, binding.JSON)
	schemaValidatorMacro(err, response)
}
