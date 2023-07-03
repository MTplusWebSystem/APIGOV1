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
 /* err := settings.SetERROR()
  if err != nil{
    Logger.Errorf("ERROR NA MAIN %v",err)
    return nil
  }*/
  router.SysRouter()
}