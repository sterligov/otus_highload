package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sterligov/otus_highload/dating/internal/domain"
	"github.com/sterligov/otus_highload/dating/internal/usecase/user"
	"go.uber.org/zap"
)

type UserHandler struct {
	userUseCase *user.UseCase
	logger      *zap.Logger
}

func NewUserHandler(uc *user.UseCase) *UserHandler {
	return &UserHandler{
		userUseCase: uc,
		logger:      zap.L().Named("user handler"),
	}
}

func (uh *UserHandler) FindByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		uh.logger.Named("FindByID").Error("strconv failed", zap.Error(err))

		c.Status(http.StatusBadRequest)
		return
	}

	_, err = uh.userUseCase.FindByID(c.Request.Context(), id)
	if err != nil {

	}

	c.JSON(http.StatusOK, gin.H{"name": "denis"})
}

func (uh *UserHandler) Create(c *gin.Context) {
	createRequest := struct {
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Sex       byte      `json:"sex"`
		Birthday  time.Time `json:"birthday"`
		CityID    int64     `json:"city_id"`
	}{}

	if err := c.BindJSON(&createRequest); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	_, err := uh.userUseCase.CreateUser(c, &domain.User{
		Email:     createRequest.Email,
		Password:  createRequest.Password,
		FirstName: createRequest.FirstName,
		LastName:  createRequest.LastName,
		Birthday:  createRequest.Birthday,
		Sex:       createRequest.Sex,
		CityID:    createRequest.CityID,
	})

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
}

func (uh *UserHandler) Filter(c *gin.Context) {
	//params := c.Request.URL.Query()
}
