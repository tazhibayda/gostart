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
	"net/http"
	_ "net/http"
)

const (
	User     = "doadmin"
	Host     = "db-postgresql-nyc3-68283-do-user-11850704-0.b.db.ondigitalocean.com"
	Database = "defaultdb"
	Password = "AVNS_aAj5qOqIwSArkVJrorC"
	Port     = 25060
	Sslmode  = "require"
)

type Man struct {
	gorm.Model
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

func getMan(context *gin.Context) {

	mans := []Man{}
	DB.Find(&mans)
	context.IndentedJSON(http.StatusOK, mans)
}

func getPerson(context *gin.Context) {
	var man Man
	DB.Where("id = ?", context.Param("id")).First(&man)
	context.IndentedJSON(http.StatusOK, man)
}

func createMan(context *gin.Context) {

	var man Man
	context.BindJSON(&man)
	DB.Create(&man)
	context.IndentedJSON(http.StatusOK, &man)
}

func updateMan(context *gin.Context) {
	var man Man
	DB.Where("id = ?", context.Param("id")).First(&man)
	context.BindJSON(&man)
	DB.Save(&man)
	context.IndentedJSON(http.StatusOK, man)
}

func deleteMan(context *gin.Context) {
	var man Man
	DB.Where("id = ?", context.Param("id")).Delete(&man)
	context.IndentedJSON(http.StatusOK, man)
}

var DB *gorm.DB
var err error

//var DSN string = fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", goes.User, goes.Password, goes.Database, goes.Host)
var url = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", User, Password, Host, Port, Database, Sslmode)

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
