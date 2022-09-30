package models

import "time"

type Karyawan struct {
	Nama 			string 		`json:"Nama" gorm:"column:Nama"`
	TanggalLahir 	time.Time 	`json:"TanggalLahir" gorm:"column:TanggalLahir"`
	JumlahAnak 		int 		`json:"JumlahAnak" gorm:"column:JumlahAnak"`
}

func (Karyawan) TableName() string {
    return "karyawan"
}