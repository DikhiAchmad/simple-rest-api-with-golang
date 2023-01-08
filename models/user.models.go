package models

import (
	"time"
)

type User struct {
    ID         uint      `json:"id" gorm:"primary_key"`
    Email	   string    `json:"email"`
    Name	   string    `json:"name"`
 	Password   string    `json:"password"`
    CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}