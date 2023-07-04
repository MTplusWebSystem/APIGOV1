package main 

import (
  "github.com/MTplusWebSystem/APIGOV1/router"
  "github.com/MTplusWebSystem/APIGOV1/settings"
  
  )
  
var (
  Logger *settings.Logger
  )

func main(){
  Logger = settings.Logs("main")
 err := settings.SetERROR()
  if err != nil{
    Logger.Error("ERROR NA MAIN",err)
    return 
  }
  router.SysRouter()
}