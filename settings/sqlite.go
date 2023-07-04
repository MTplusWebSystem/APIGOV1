package settings 

import (
  "gorm.io/gorm"
  )

func SetSQLite()(*gorm.DB,error){
  logger := Logs("sqlite")
  
}