package middleware

import (
	"encoding/json"
	"time"

	"github.com/sterligov/otus_highload/dating/internal/server/http/handler"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/sterligov/otus_highload/dating/internal/config"
	"github.com/sterligov/otus_highload/dating/internal/domain"
	"github.com/sterligov/otus_highload/dating/internal/usecase/auth"
	"go.uber.org/zap"
)

const identityKey = "email"

func Auth(cfg *config.Config, auth *auth.UseCase) *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(cfg.JWT.SecretKey),
		Timeout:     24 * time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*domain.User); ok {
				v.Password = ""
				return jwt.MapClaims{
					identityKey: v.Email,
					"user":      v,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			u, err := GetUserFromClaims(c)
			if err != nil {
				return nil
			}

			return u
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			loginRequest := struct {
				Email    string `json:"email"`
				Password string `json:"password"`
			}{}
			if err := c.ShouldBind(&loginRequest); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			u, err := auth.Login(c, loginRequest.Email, loginRequest.Password)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return u, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			_, ok := data.(*domain.User)

			return ok
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		zap.L().Named("auth middleware").Error("jwt failed", zap.Error(err))
	}

	err = authMiddleware.MiddlewareInit()
	if err != nil {
		zap.L().Named("auth middleware").Error("MiddlewareInit failed", zap.Error(err))
	}

	return authMiddleware
}

func GetUserFromClaims(c *gin.Context) (*domain.User, error) {
	claims := jwt.ExtractClaims(c)
	user, ok := claims["user"]
	if !ok {
		return nil, handler.ErrBadClaims
	}

	ju, err := json.Marshal(user)
	if err != nil {
		return nil, handler.ErrBadClaims
	}

	u := new(domain.User)
	if err := json.Unmarshal(ju, u); err != nil {
		return nil, handler.ErrBadClaims
	}

	return u, nil
}
