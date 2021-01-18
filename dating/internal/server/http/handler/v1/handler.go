package v1

import (
	"net/http"

	"github.com/sterligov/otus_highload/dating/internal/server/http/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
)

func NewHandler(
	uh *UserHandler,
	ch *CityHandler,
	jm *jwt.GinJWTMiddleware,
) http.Handler {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	v := r.Group("v1")
	{
		v.POST("/sign-in", jm.LoginHandler)
		v.POST("/sign-out", jm.LogoutHandler)
		v.POST("/sign-up", uh.Create)
		v.GET("/cities", ch.FindAll)
		v.Use(jm.MiddlewareFunc())
		{
			v.GET("/users", uh.FindAll)
			v.GET("/users/:id", uh.FindByID)
			v.GET("/users/:id/friends", uh.FindFriends)
			v.POST("/friends/:friend_id", uh.Subscribe)
			v.DELETE("/friends/:friend_id", uh.Unsubscribe)
		}
	}

	return r
}
