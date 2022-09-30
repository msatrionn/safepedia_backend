package routes

import (
	controllers "backend/Controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
    r := gin.Default()
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"PUT", "PATCH"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
          return origin == "https://github.com"
        },
        MaxAge: 12 * time.Hour,
      }))
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })
    r.GET("/karyawan", controllers.GetKaryawan)
    r.GET("/karyawan/search", controllers.GetKaryawanWhere)
    r.GET("/kendaraan", controllers.GetKendaraanPerusahaan)
    r.GET("/kendaraan/search", controllers.GetKendaraanWhere)
    r.GET("/kendaraan/jenis", controllers.GetKendaraanType)
    return r
}