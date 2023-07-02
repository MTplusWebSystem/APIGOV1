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
		
	v1.POST("opening",func(c *gin.Context) {
		c.JSON(200, gin.H{
			"opening": "POST 200",
		})})
	
	v1.DELETE("opening",func(c *gin.Context) {
		c.JSON(200, gin.H{
			"opening": "DELETE 200",
		})})
		
	v1.PUT("opening",func(c *gin.Context) {
		c.JSON(200, gin.H{
			"opening": "PUT 200",
		})})
		
	v1.GET("openings",)
}