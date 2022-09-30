package controllers

import (
	models "backend/Models"
	"strconv"

	pagination "github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetKaryawan(c *gin.Context) {
    pages := c.Query("page")
    limits := c.Query("limit")

    page, _ := strconv.Atoi(c.DefaultQuery("page", pages))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", limits))
    var karyawan []models.Karyawan
    db := c.MustGet("db").(*gorm.DB)

    paginator := pagination.Paging(&pagination.Param{
        DB:      db,
        Page:    page,
        Limit:   limit,
        OrderBy: []string{"id desc"},
        ShowSQL: true,
    }, &karyawan)
    c.JSON(200, paginator)
}
func GetKaryawanWhere(c *gin.Context) {
    nama := c.Query("nama")
    tgl := c.Query("lahir")
    jml := c.Query("jml_anak")
    var karyawan []models.Karyawan
    db := c.MustGet("db").(*gorm.DB)
    db.Where("Nama LIKE ?", "%"+ nama +"%").Or("TanggalLahir LIKE ?", "%"+tgl+"%").Or("JumlahAnak LIKE ?", "%"+jml+"%").Find(&karyawan)
    c.JSON(200, karyawan)
}

