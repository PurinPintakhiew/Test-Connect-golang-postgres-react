package api

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/gofiber/fiber/v2"
)

type data struct {
	Id   int64 `json:"id"`
	Name string `json:"name"`
}

var DB *gorm.DB

func SetupPostgres() {
	dsn := "host=10.133.0.131 user=purin password=password dbname=list port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if db != nil{
		println("connect")
		DB = db
		db.AutoMigrate(&data{})
	}
}

func DataList(c *fiber.Ctx) error {
	var list []data
	DB.Find(&list)
	return c.JSON(list)
}