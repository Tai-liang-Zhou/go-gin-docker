package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tsmc.com/go-gin-docker/domain"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

type MariaRouter struct {
	MariaUsecase domain.MariaUsecase
}

func NewMariaRouter(r *gin.Engine, mariaUsecase domain.MariaUsecase) {
	router := &MariaRouter{
		MariaUsecase: mariaUsecase,
	}

	r.GET("/api/v1/customers", router.GetCustomers)
}

func (m *MariaRouter) GetCustomers(c *gin.Context) {
	res, err := m.MariaUsecase.GetCustomers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}
	c.JSON(http.StatusOK, res)
}
