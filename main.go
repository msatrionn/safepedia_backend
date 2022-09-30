package main

// only need mysql OR sqlite
// both are included here for reference
import (
	"backend/routes"

	config "backend/config"

	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db := config.SetupDB()
    // db.AutoMigrate(&models.Karyawan{})
    r := routes.SetupRoutes(db)
 	r.Run(":8080")
}