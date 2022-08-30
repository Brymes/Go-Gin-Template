package api

//Code for JWT implementation
/* 
import (
	config "Core-Wallet-Service/config"
	"Core-Wallet-Service/models"
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

// MAP Key used in extracting claims
var identityKey = "user_id"

func JwtMiddleWareConfig() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour * 168,
		IdentityKey: identityKey,
		IdentityHandler: func(c *gin.Context) interface{} {
			var usr models.User

			claims := jwt.ExtractClaims(c)
			mongoId, _ := primitive.ObjectIDFromHex(claims[identityKey].(string))

			err := config.MongoClient.Collection("Users").FindOne(context.TODO(), bson.M{"_id": mongoId}).Decode(&usr)
			if err != nil {
				return nil
			}

			return &usr
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
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware

}
 */