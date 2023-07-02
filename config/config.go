package config

import "gorm.io/gorm"

var(
  db *gorm.DB
  logger *Logger 
  )

func StartDB () error{
  return nil
}

func ResponseLogger(p string) *Logger{
  logger = NewLogger(p)
  return logger
}