package models

import "time"

type User struct {
	Id        uint      `gorm:"primarykey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"-"`
	Telp      string    `json:"telp"`
	Role      string    `json:"role"`
	Nik       string    `json:"nik"`
	Level     string    `json:"level"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
