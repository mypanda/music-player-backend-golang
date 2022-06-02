package v2

import (
	"music-player-backend/response"

	"github.com/gin-gonic/gin"
)

type GoodsController struct{}

func NewGoodsController() GoodsController {
	return GoodsController{}
}

// v2 商品详情
func (g *GoodsController) GoodsOne(c *gin.Context) {
	response.Success("v2 one")
}

// v2 商品列表
func (g *GoodsController) GoodsList(c *gin.Context) {
	response.Success("v2 list")
}
