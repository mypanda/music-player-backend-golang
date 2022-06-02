package route

import (
	"log"
	"music-player-backend/controller"
	v1 "music-player-backend/controller/v1"
	v2 "music-player-backend/controller/v2"
	"music-player-backend/db"
	"music-player-backend/response"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router() *gin.Engine {
	var DB *mongo.Client = db.ConnectDB()

	r := gin.Default()

	r.Use(response.SetContext)
	//
	r.NoRoute(HandleNotFound)
	r.NoMethod(HandleNotFound)
	r.Use(Recover)

	//无group
	indexc := controller.NewIndexController()
	r.GET("/index/index", indexc.Index)

	//v1 group
	apiv1 := r.Group("/v1")
	{
		goodsc := v1.NewGoodsController()
		apiv1.GET("/goods/one", goodsc.GoodsOne)
		apiv1.GET("/goods/list", goodsc.GoodsList)
	}
	{
		testc := v1.NewTestController()
		apiv1.GET("/test", testc.GetList(DB))
		apiv1.GET("/test/:id", testc.Get(DB))
		apiv1.PUT("/test", testc.Update(DB))
		apiv1.POST("/test", testc.Create(DB))
		apiv1.DELETE("/test", testc.Delete(DB))
	}
	{
		uploadc := v1.NewUploadController()
		apiv1.POST("/upload", uploadc.Create(DB))
	}
	//v2 group
	apiv2 := r.Group("/v2")
	{
		goodsc := v2.NewGoodsController()
		apiv2.GET("/goods/one", goodsc.GoodsOne)
		apiv2.GET("/goods/list", goodsc.GoodsList)
	}
	return r
}

func HandleNotFound(c *gin.Context) {
	response.Error(404, "资源未找到")
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			response.Error(500, "服务器内部错误")
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}
