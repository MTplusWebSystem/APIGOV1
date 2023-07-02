package router

import "github.com/gin-gonic/gin"
func Start(){
  msg()
  r := gin.Default()
	start_control(r)
	r.Run()
}