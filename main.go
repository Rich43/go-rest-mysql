package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type singles struct {
	gorm.Model
	Id        int    `json:"ID" gorm:"primary_key"`
	Title     string `json:"title"`
	Artist    string `json:"artist"`
	Lyrics    string ` json:"lyrics"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
	DeletedAt string `json:"DeletedAt"`
}

var DB *gorm.DB

func getSingles(c *gin.Context) {
	var singlesCollection []singles
	DB.Create(&singlesCollection)
	c.JSON(http.StatusOK, singlesCollection)
}

func addSingle(c *gin.Context) {
	var singlesCollection = []singles{
		{Title: c.GetString("title"), Artist: c.GetString("artist"), Lyrics: c.GetString("lyrics")},
	}
	DB.Create(&singlesCollection)
	c.JSON(http.StatusOK, singlesCollection)
}

func main() {
	db, err := gorm.Open(mysql.Open("root:Password1@tcp(127.0.0.1:3306)/musicdb"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	migrateErr := db.AutoMigrate(&singles{})
	if migrateErr != nil {
		panic("failed to migrate")
	}
	DB = db
	router := gin.Default()
	context := gin.Context{}
	context.Set("db", db)
	router.GET("/api/singles", getSingles)
	router.POST("/api/singles", addSingle)
	routerErr := router.Run("localhost:8080")
	if routerErr != nil {
		panic("failed to start server")
	}
}
