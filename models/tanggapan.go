package models

import "time"

type Tanggapan struct {
	Id            uint      `gorm:"primarykey" json:"id"`
	PengaduanId   int       `json:"pengaduan_id"`
	Pengaduan     Pengaduan `json:"-" gorm:"foreignKey:PengaduanId"`
	Tgl_Tanggapan string    `json:"tgl_tanggapan"`
	Tanggapan     string    `json:"tanggapan"`
	UserId        int       `json:"user_id"`
	User          User      `json:"-" gorm:"foreignKey:UserId"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}
