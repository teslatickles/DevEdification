package controllers

import (
	"github.com/DevEdification/v2/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}


var identityKey = "id"

func HandleLogin() *jwt.GinJWTMiddleware {

	authware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "test zone",
		SigningAlgorithm: "SHA256",
		Key:              []byte("secret key"),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			//var loginValues login
			if err := c.ShouldBind(&models.User{
				Username: "test",
				Password: "test",
				Email:    "test",
				Role:     "test",
			}); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := models.User{}.Username
			password := models.User{}.Password

			if userID == "admin" && password == "admin" || userID == "test" && password == "test" {
				return &models.User{
					Username: userID,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.User); ok && v.Username == "admin" {
				return true
			}
			return false
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		LoginResponse:   nil,
		LogoutResponse:  nil,
		RefreshResponse: nil,
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Username: claims[identityKey].(string),
			}
		},
		IdentityKey: identityKey,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",
		TokenHeadName:         "Bearer",
		TimeFunc:              time.Now,
		HTTPStatusMessageFunc: nil,
		PrivKeyFile:           "",
		PubKeyFile:            "",
		SendCookie:            false,
		CookieMaxAge:          0,
		SecureCookie:          false,
		CookieHTTPOnly:        false,
		CookieDomain:          "",
		SendAuthorization:     false,
		DisabledAbort:         false,
		CookieName:            "dadsbreakfast",
		CookieSameSite:        0,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authware
}

func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*models.User).Username,
		"text":     "Hello World Domination",
	})
}
