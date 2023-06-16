package initRouter

import (
	"MServer/handler"
	"MServer/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	//添加路由
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/", handler.Index)

	//router.POST("/api/Login", handler.Login)

	//g1 := router.Group("/v1").Use(middleware1)
	//By Group設定middleware
	//g1.GET("/getting", func(c *gin.Context) {
	//	fmt.Println("doing v1 getting")
	//	c.JSON(http.StatusOK, gin.H{"data": "v1 getting"})
	//})

	//grp1 := router.Group("/User")
	//{
	//	grp1.GET("user", handler.GetUsers)
	//	grp1.POST("user", handler.CreateUser)
	//	grp1.GET("user/:id", handler.GetUserByID)
	//	//grp1.PUT("user/:id", handler.UpdateUser)
	//	//grp1.DELETE("user/:id", handler.DeleteUser)
	//}

	grpAPI := router.Group("/api")
	{
		grpAPI.GET("MaintainItemList", handler.GetMaintainItemList)
		grpAPI.GET("MaintainItemList/:ft", handler.GetMaintainItemListByFT)
		grpAPI.POST("CreateMaintainOrder", handler.CreateMaintainOrder)
		grpAPI.GET("MaintainOrderList", handler.GetMaintainOrderAll)
		//grpAPI.GET("MaintainOrderSolvedList", handler.GetMaintainOrderSolvedList)
		//grpAPI.GET("MaintainOrderHandlingList", handler.GetMaintainOrderHandlingList)
		//grpAPI.GET("MaintainOrder/:id", handler.GetMaintainOrderById)
	}

	return router
}
