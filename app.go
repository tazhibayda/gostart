package main

import (
	_ "encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	_ "net/http"
)

var DB *gorm.DB
var err error

var DSN string = fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", user, password, database, host)
var url = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", user, password, host, port, database, sslmode)

func createDb() {

	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Error connecting db")
	}
	DB.AutoMigrate(&Man{})

}

func initRouter() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/list", getMan)
	router.GET("/list/:id", getPerson)
	router.POST("/create", createMan)
	router.PUT("/list/:id", updateMan)
	router.DELETE("/list/:id", deleteMan)

	log.Fatal(router.Run())
}

func main() {

	createDb()
	initRouter()
	// this
}
