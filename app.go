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
	user     = "doadmin"
	host     = "db-postgresql-nyc3-68283-do-user-11850704-0.b.db.ondigitalocean.com"
	database = "defaultdb"
	password = "AVNS_aAj5qOqIwSArkVJrorC"
	port     = 25060
	sslmode  = "require"
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

func initRouter() {
	router := gin.Default()
	router.Use(cors.Default())
	//gin.HandlerFunc("/list", getMan)
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
