package controllers

import (
	models "backend/Models"
	"strconv"

	pagination "github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetKendaraanPerusahaan(c *gin.Context) {
    pages := c.Query("page")

    page, _ := strconv.Atoi(c.DefaultQuery("page", pages))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "1"))
    var Kendaraan []models.KendaraanPerusahaan
    db := c.MustGet("db").(*gorm.DB)

    paginator := pagination.Paging(&pagination.Param{
        DB:      db,
        Page:    page,
        Limit:   limit,
        OrderBy: []string{"id desc"},
        ShowSQL: true,
    }, &Kendaraan)
    c.JSON(200, paginator)
}

func GetKendaraanWhere(c *gin.Context) {
    tipe := c.Query("tipe")
    warna := c.Query("warna")
    tgl := c.Query("tgl")
    jml := c.Query("jml")
    var KendaraanPerusahaan []models.KendaraanPerusahaan
    db := c.MustGet("db").(*gorm.DB)
    db.Where("Tipe LIKE ?", "%"+ tipe +"%").Or("warna LIKE ?", "%"+warna+"%").Or("TanggalBeli LIKE ?", "%"+tgl+"%").Or("JumlahRoda LIKE ?", "%"+jml+"%").Find(&KendaraanPerusahaan)
    c.JSON(200, KendaraanPerusahaan)
}

func GetKendaraanType(c *gin.Context) {
    tipe := c.Query("tipe")
    var KendaraanPerusahaan []models.KendaraanPerusahaan
    db := c.MustGet("db").(*gorm.DB)
    db.Where("Tipe LIKE ?", "%"+ tipe +"%").Find(&KendaraanPerusahaan)
    c.JSON(200, KendaraanPerusahaan)
}


