package schemes

import (
  "gorm.io/gorm")

type Openings struct {
  gorm.Model
  Role  string
  CompanyName string 
  Location string
  Remote bool
  Link string
  Salary int64
}