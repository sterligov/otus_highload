package v1

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
)

func NewHandler(uh *UserHandler, jm *jwt.GinJWTMiddleware) http.Handler {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v := r.Group("v1")
	{
		v.POST("/login", jm.LoginHandler)
		v.POST("/register", uh.Create)
		v.Use(jm.MiddlewareFunc())
		{
			v.GET("/users/:id", uh.FindByID)
			v.GET("/users/filter", uh.Filter)
		}
	}

	return r
}
