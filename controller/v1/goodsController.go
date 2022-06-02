package v1

import (
	"music-player-backend/response"

	"github.com/gin-gonic/gin"
)

type GoodsController struct{}

func NewGoodsController() GoodsController {
	return GoodsController{}
}

// v1 商品详情
func (g *GoodsController) GoodsOne(c *gin.Context) {
	response.Success("v1 one")
}

// v1 商品列表
func (g *GoodsController) GoodsList(c *gin.Context) {
	response.Success("v1 one")
}
