package router

import "github.com/gin-gonic/gin"

func SysRouter(){
  system := gin.Default()
  endpoint(system)
  system.Run()
}