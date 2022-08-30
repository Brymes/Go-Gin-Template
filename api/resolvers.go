package api

import (	
	"github.com/gin-gonic/gin"
)

func EndPointName(c *gin.Context) {
	// var request ReqType

	//Bind Requst Payload to request type
	// if err := c.BindJSON(&request); err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, err)
	// }

	// reqBuffer, reqLogger := config.InitRequestLogger(request.Service)

	//Call Controller
	// response := request.CreateUploadURL(serverURL, reqLogger)
	
	// Log request logs to Central Logs
	// defer log.Println(reqBuffer)
	
	//Return Response
	/* if response["status"].(bool) != true {
		c.IndentedJSON(http.StatusBadRequest, response)
	} else {
		c.IndentedJSON(http.StatusOK, response)
	} */
}
