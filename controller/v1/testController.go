package v1

import (
	"music-player-backend/response"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type TestForm struct {
	Name   string `json:"name" binding:"required"`
	Singer string `json:"singer" binding:"required"`
	Url    string `json:"url" binding:"required"`
}

type TestController struct{}

func NewTestController() TestController {
	return TestController{}
}

func (g *TestController) Get(DB *mongo.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		response.Success("v1 Get " + id)
	}
}
func (g *TestController) Create(DB *mongo.Client) func(c *gin.Context) {
	return func(c *gin.Context) {

		testReq := TestForm{}
		if err := c.Bind(&testReq); err != nil {
			response.Error(200001, err.Error())
			return
		}

		// data := map[string]interface{}{
		// 	"name": "panda",
		// }

		response.Success(testReq)
	}
}
func (g *TestController) Delete(DB *mongo.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		response.Success("v1 Delete")
	}
}

func (g *TestController) GetList(DB *mongo.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		response.Success("v1 GetList")
	}
}
func (g *TestController) Update(DB *mongo.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		response.Success("v1 Update")
	}
}
