package controllers 

import "github.com/gin-gonic/gin"
func CreateHandle(c *gin.Context) {
  c.JSON(200, gin.H{
			"opening": "POST 200",})
}

func UpdateHandle(c *gin.Context) {
  c.JSON(200, gin.H{
			"opening": "PUT 200",})
}

func DeleteHandle(c *gin.Context) {
  c.JSON(200, gin.H{
			"opening": "DELETE 200",})
}
func ShowHandle(c *gin.Context) {
  c.JSON(200, gin.H{
			"opening": "GET 200",})
}

func AllistHandle(c *gin.Context) {
  c.JSON(200, gin.H{
			"opening": "GET 200",})
}