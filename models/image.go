package models

import "time"

type Image struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	PengaduanID int       `json:"pengaduan_id"`
	Pengaduan   Pengaduan `json:"-" gorm:"foreignKey:PengaduanID"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
