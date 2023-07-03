package router

import "github.com/gin-gonic/gin"
import "github.com/MTplusWebSystem/APIGOV1/controllers"

func endpoint(controller *gin.Engine){
  v1 := controller.Group("/api/v1")
  v1.POST("opening",controllers.CreateHandle)
	v1.PUT("opening",controllers.UpdateHandle)
	v1.DELETE("opening",controllers.DeleteHandle)
	v1.GET("openings",controllers.AllistHandle)
	v1.GET("opening",controllers.ShowHandle)
}