package router
import "github.com/gin-gonic/gin"
func msg(){
  yellow := "\003[1;33m"
  print(yellow+"SERVIDOR INICIADO NA PORTA 8080")
}
func start_control(controller *gin.Engine){
  v1 := controller.Group("/api/v1")
  v1.GET("opening",func(c *gin.Context) {
		c.JSON(200, gin.H{
			"opening": "GET 200",
		})})
}