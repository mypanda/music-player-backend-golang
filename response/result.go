package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	r *gin.Context
)

func SetContext(ctx *gin.Context) {
	r = ctx
}

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) {
	res := Result{0, "success", data}
	r.JSON(http.StatusOK, res)
}

func Error(code int, msg string) {
	res := Result{code, msg, gin.H{}}
	r.JSON(http.StatusOK, res)
	r.Abort()
}
