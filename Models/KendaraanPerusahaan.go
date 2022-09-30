package models

import "time"

type KendaraanPerusahaan struct {
	Tipe 		string 		`json:"Tipe" gorm:"column:Tipe"`
	Warna 		string 		`json:"Warna" gorm:"column:Warna"`
	TanggalBeli time.Time 	`json:"TanggalBeli" gorm:"column:TanggalBeli"`
	JumlahRoda 	int 		`json:"JumlahRoda" gorm:"column:JumlahRoda"`
}

func (KendaraanPerusahaan) TableName() string {
    return "KendaraanPerusahaan"
}