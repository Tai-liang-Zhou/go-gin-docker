package routers

import (
	"testing"

	"github.com/gin-gonic/gin"
	"tsmc.com/go-gin-docker/domain"
)

func TestMariaRouter_GetCustomers(t *testing.T) {
	type fields struct {
		MariaUsecase domain.MariaUsecase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MariaRouter{
				MariaUsecase: tt.fields.MariaUsecase,
			}
			m.GetCustomers(tt.args.c)
		})
	}
}
