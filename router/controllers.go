package router
import "github.com/gin-gonic/gin"
import "github.com/MTplusWebSystem/APIGOV1/handler"


func start_control(controller *gin.Engine){
  v1 := controller.Group("/api/v1")
  
  v1.POST("opening",handler.CreateHandle)
	v1.PUT("opening",handler.UpdateHandle)
	v1.DELETE("opening",handler.DeleteHandle)
	v1.GET("openings",handler.AllistHandle)
	v1.GET("opening",handler.ShowHandle)
}