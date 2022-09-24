package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tsmc.com/go-gin-docker/domain"
	e "tsmc.com/go-gin-docker/pkg"
	"tsmc.com/go-gin-docker/pkg/httputil"
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

// ListAccounts godoc
// @Summary      List Customers
// @Description  get Customers
// @Tags         Customers
// @Produce      json
// @Success      200  {object}  domain.Customers
// @Router       /api/v1/customers [get]
func (m *MariaRouter) GetCustomers(c *gin.Context) {
	r := httputil.NewResponse{C: c}
	res, err := m.MariaUsecase.GetCustomers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}
	r.Response(http.StatusOK, e.SUCCESS, res)
}

// ListAccounts godoc
// @Summary      List Customers
// @Description  get Customers
// @Tags         Customers
// @Produce      json
// @Success      200  {object}  domain.Customers
// @Router       /api/v1/customers/:id [get]
func (m *MariaRouter) GetCustomerbyID(c *gin.Context) {
	r := httputil.NewResponse{C: c}


	res, err := m.MariaUsecase.GetCustomers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}
	r.Response(http.StatusOK, e.SUCCESS, res)
}
