package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
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
