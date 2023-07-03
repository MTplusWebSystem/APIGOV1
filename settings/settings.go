package settings

import (
  "gorm.io/gorm"
  )
  
var(
  db *gorm.DB
  logger *Logger 
  )
  
func SetERROR() error{
  return nil
}

func Logs(p string) *Logger{
  logger = NewLogger(p)
  return logger
}