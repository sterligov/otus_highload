package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sterligov/otus_highload/dating/internal/domain"
	"github.com/sterligov/otus_highload/dating/internal/server/http/handler"
	"github.com/sterligov/otus_highload/dating/internal/usecase/city"
	"go.uber.org/zap"
)

type (
	City struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	CityResponse struct {
		City
		Country `json:"country"`
	}

	Country struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	CityHandler struct {
		cityUseCase *city.UseCase
		logger      *zap.Logger
	}
)

func NewCityHandler(cu *city.UseCase) *CityHandler {
	return &CityHandler{
		cityUseCase: cu,
		logger:      zap.L().Named("city handler"),
	}
}

func (uh *CityHandler) FindAll(c *gin.Context) {
	cities, err := uh.cityUseCase.FindAll(c)
	if err != nil {
		handler.JSONError(c, err)
		return
	}

	c.JSON(http.StatusOK, ToCities(cities))
}

func ToCities(domainCities []*domain.City) []*CityResponse {
	cities := make([]*CityResponse, 0, len(domainCities))

	for _, c := range domainCities {
		cities = append(cities, ToCity(c))
	}

	return cities
}

func ToCity(c *domain.City) *CityResponse {
	return &CityResponse{
		City: City{
			ID:   c.ID,
			Name: c.Name,
		},
		Country: Country{
			ID:   c.Country.ID,
			Name: c.Country.Name,
		},
	}
}
