package httputil

import (
	"github.com/gin-gonic/gin"
	_ "tsmc.com/go-gin-docker/pkg"
	e "tsmc.com/go-gin-docker/pkg"
)

type NewResponse struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *NewResponse) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
}
