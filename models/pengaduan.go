package models

import "time"

type Pengaduan struct {
	Id            uint      `gorm:"primarykey" json:"id"`
	Tgl_Pengaduan string    `json:"tgl_pengaduan"`
	JudulLaporan  string    `json:"judul_laporan"`
	Isi_Laporan   string    `json:"isi_laporan"`
	Status        string    `json:"status"`
	UserId        int       `json:"user_id"`
	User          User      `json:"-" gorm:"foreignKey:UserId"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}
