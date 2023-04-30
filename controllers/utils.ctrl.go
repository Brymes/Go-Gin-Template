package controllers

/*
	var Levels = map[string]bool{
		"ADMIN": true, "USER": true,
	}

	func authMacro(c *gin.Context, level string, response *schemas.Response) (user *models.User) {
		var res bool

		usr, _ := c.Get(IdentityKey)
		user = usr.(*models.User)

		switch level {
		case "ALL":
			_, res = Levels[user.Role]
		case user.Role:
			res = true
		default:
			res = false
		}
		if !res {
			response.Status, response.Message = http.StatusUnauthorized, "You do not have access to this endpoint"
			panic("Unauthorized")
		}
		return user
	}
*/
/*func validatorMacro(c *gin.Context, validator interface{}, response *schemas.Response) {
	//Bind Request Payload to request type
	if err := c.ShouldBindBodyWith(&validator, binding.JSON); err != nil {
		response.Status, response.Message = 400, "Invalid Request Payload"+err.Error()
		panic("Validation failed")
	}
}
*/
