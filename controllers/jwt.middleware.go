package controllers

import (
	"App-Name/config"
	"App-Name/models"
	"App-Name/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

// MAP Key used in extracting claims
var IdentityKey = "user_id"

func JwtMiddleWareConfig() *jwt.GinJWTMiddleware {
	buff, logger := config.InitRequestLogger("JWT")

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(os.Getenv("SECRET_KEY")),
		Timeout:     168 * time.Hour,
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			var usr models.User

			claims := jwt.ExtractClaims(c)

			result := config.DBClient.Model(&models.User{}).Where("email = ?", claims[IdentityKey]).First(&usr)
			_, code := utils.HandleDbError(result, logger)
			if code != 0 {
				logger.Println(result.Error)
				return nil
			}

			return &usr
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			defer log.Println(buff)
			var loginVals map[string]string
			var usr models.User

			// Does this work??
			if err := c.ShouldBind(&loginVals); err != nil {
				log.Println("na ham")
				return "", jwt.ErrMissingLoginValues
			}

			email := loginVals["email"]
			password := loginVals["password"]

			result := config.DBClient.Model(&models.User{}).Where("email = ?", email).First(&usr)
			_, code := utils.HandleDbError(result, logger)
			if code != 0 {
				logger.Println(result.Error)
				return nil, jwt.ErrFailedAuthentication
			}

			if utils.CheckPasswordHash(password, usr.Password) {
				return &usr, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			logger.Println(message)
			logger.Println(c.GetHeader("Authorization"))
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		logger.Fatal("JWT Error:" + err.Error())
	}
	log.Println(buff)

	return authMiddleware

}
