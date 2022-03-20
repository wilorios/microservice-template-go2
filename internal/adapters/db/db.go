package db

import (
	"fmt"
	"log"

	"github.com/wilorios/microservice-template-go2/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func Init() *gorm.DB {

	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432"
	DBConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	DBConn.AutoMigrate(&entity.Person{})
	fmt.Println("started database")

	return DBConn
}
