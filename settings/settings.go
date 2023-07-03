package settings

import (
  "gorm.io/gorm"
  )
  
var(
  db *gorm.DB
  logger *Logger 
  )
  
func Set(){
  return nil
}