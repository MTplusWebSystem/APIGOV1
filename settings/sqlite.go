package settings 
/*
import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
  )

type Company struct {
  gorm.Model
  Code  string
  Price uint
}

func SetSQLite()(*gorm.DB,error){
  logger := Logs("sqlite")
  dbPath:= "./db/main.db"
  
  
  db, err:= gorm.Open(sqlite.Open(dbPath,&gorm.Config{}))
  
  if err !=nil{
    logger.Errof("SQLITE ERRO %v",erro)
    return nil, err
  }
}*/