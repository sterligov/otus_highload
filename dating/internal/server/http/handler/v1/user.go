package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/sterligov/otus_highload/dating/internal/server/http/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sterligov/otus_highload/dating/internal/domain"
	"github.com/sterligov/otus_highload/dating/internal/server/http/handler"
	"github.com/sterligov/otus_highload/dating/internal/usecase/user"
	"go.uber.org/zap"
)

const dateLayout = "2006-01-02"

type (
	User struct {
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Birthday  string `json:"birthday"`
		Interests string `json:"interests"`
		Sex       string `json:"sex"`
	}

	UserRequest struct {
		CityID   int64  `json:"city_id"`
		Password string `json:"password"`
		User
	}

	UserResponse struct {
		ID   int64 `json:"id"`
		City *City `json:"city"`
		User
	}

	UserHandler struct {
		userUseCase *user.UseCase
		logger      *zap.Logger
	}
)

func NewUserHandler(uc *user.UseCase) *UserHandler {
	return &UserHandler{
		userUseCase: uc,
		logger:      zap.L().Named("user handler"),
	}
}

func (uh *UserHandler) FindAll(c *gin.Context) {
	users, err := uh.userUseCase.FindAll(c)
	if err != nil {
		uh.logger.Error("user find all failed", zap.Error(err))

		handler.JSONError(c, err)
		return
	}

	c.JSON(http.StatusOK, ToUsersResponse(users))
}

func (uh *UserHandler) FindByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		uh.logger.Error("strconv failed", zap.Error(err))

		handler.JSONError(c, handler.ErrBadParam)
		return
	}

	u, err := uh.userUseCase.FindByID(c, id)
	if err != nil {
		uh.logger.Error("FindByID failed", zap.Error(err))

		handler.JSONError(c, err)
		return
	}

	c.JSON(http.StatusOK, ToUserResponse(u))
}

func (uh *UserHandler) Subscribe(c *gin.Context) {
	curUser, err := middleware.GetUserFromClaims(c)
	if err != nil {
		uh.logger.Warn("get user from claims failed", zap.Error(err))

		handler.JSONError(c, handler.ErrBadParam)
		return
	}

	friendID, err := strconv.ParseInt(c.Param("friend_id"), 10, 64)
	if err != nil {
		uh.logger.Error("strconv failed", zap.Error(err))

		handler.JSONError(c, handler.ErrBadParam)
		return
	}

	insertedID, err := uh.userUseCase.Subscribe(c, curUser.ID, friendID)
	if err != nil {
		uh.logger.Error("Subscribe failed", zap.Error(err))

		handler.JSONError(c, err)
		return
	}

	c.JSON(http.StatusOK, Inserted{InsertedID: insertedID})
}

func (uh *UserHandler) Unsubscribe(c *gin.Context) {
	curUser, err := middleware.GetUserFromClaims(c)
	if err != nil {
		uh.logger.Warn("get user from claims failed", zap.Error(err))

		handler.JSONError(c, handler.ErrBadParam)
		return
	}

	friendID, err := strconv.ParseInt(c.Param("friend_id"), 10, 64)
	if err != nil {
		uh.logger.Error("strconv failed", zap.Error(err))

		handler.JSONError(c, handler.ErrBadParam)
		return
	}

	affected, err := uh.userUseCase.Unsubscribe(c, curUser.ID, friendID)
	if err != nil {
		uh.logger.Error("Subscribe failed", zap.Error(err))

		handler.JSONError(c, err)
		return
	}

	c.JSON(http.StatusOK, Affected{Affected: affected})
}

func (uh *UserHandler) FindFriends(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		uh.logger.Error("strconv failed", zap.Error(err))

		handler.JSONError(c, handler.ErrBadParam)
		return
	}

	users, err := uh.userUseCase.FindFriends(c, id)
	if err != nil {
		uh.logger.Error("subscribe failed", zap.Error(err))

		handler.JSONError(c, err)
		return
	}

	c.JSON(http.StatusOK, ToUsersResponse(users))
}

func (uh *UserHandler) Create(c *gin.Context) {
	request := new(UserRequest)

	if err := c.BindJSON(&request); err != nil {
		uh.logger.Error("bind json failed", zap.Error(err))

		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad JSON"})
		return
	}

	birthday, err := time.Parse(dateLayout, request.Birthday)
	if err != nil {
		uh.logger.Error("time parse", zap.Error(err))

		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad JSON"})
		return
	}

	insertedID, err := uh.userUseCase.CreateUser(c, &domain.User{
		Email:     request.Email,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Interests: request.Interests,
		Birthday:  birthday,
		Sex:       request.Sex,
		City: &domain.City{
			ID: request.CityID,
		},
	})

	if err != nil {
		uh.logger.Error("create user failed", zap.Error(err))

		handler.JSONError(c, err)
		return
	}

	c.JSON(http.StatusOK, Inserted{InsertedID: insertedID})
}

func ToUsersResponse(domainUsers []*domain.User) []*UserResponse {
	users := make([]*UserResponse, 0, len(domainUsers))

	for _, u := range domainUsers {
		users = append(users, ToUserResponse(u))
	}

	return users
}

func ToUserResponse(u *domain.User) *UserResponse {
	return &UserResponse{
		ID: u.ID,
		User: User{
			Email:     u.Email,
			Interests: u.Interests,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Birthday:  u.Birthday.Format(dateLayout),
			Sex:       u.Sex,
		},
		City: &City{
			ID:   u.City.ID,
			Name: u.City.Name,
		},
	}
}