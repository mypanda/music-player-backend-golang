package controller

import (
	"fmt"
	"music-player-backend/response"
	"time"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func NewIndexController() IndexController {
	return IndexController{}
}

//首页，返回一个成功的提示
func (g *IndexController) Index(c *gin.Context) {
	fmt.Println("controller:index: " + time.Now().String())
	response.Success("success")
}
