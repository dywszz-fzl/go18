package main

import (
	"awesomeProject/book/v2/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type Book struct {
	ID     uint    `json:"id" gorm:"primarykey"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

func setupDatabase() *gorm.DB {
	mc := config.C().MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mc.Username,
		mc.Password,
		mc.Host,
		mc.Port,
		mc.DB)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Book{})
	return db
}

func main() {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "application.yaml"
	}

	config.LoadConfigfromYaml(path)

	r := gin.Default()
	db := setupDatabase()

	r.POST("/books", func(c *gin.Context) {
		var book Book
		if err := c.ShouldBind(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&book)
		c.JSON(http.StatusCreated, gin.H{"data": book})
	})

	r.GET("/books", func(c *gin.Context) {
		var books []Book
		db.Find(&books)
		c.JSON(http.StatusOK, gin.H{"data": books})

	})

	r.GET("/books/:id", func(c *gin.Context) {
		var book Book
		id := c.Params.ByName("id")
		if err := db.First(&book, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"data": book})
	})

	r.PUT("/books/:id", func(c *gin.Context) {
		var book Book
		id := c.Params.ByName("id")
		if err := db.First(&book, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book Not Found"})
			return
		}

		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Save(&book)
		c.JSON(http.StatusOK, book)
	})

	r.DELETE("/books/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		if err := db.Delete(&Book{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	})

	ac := config.C().Application
	r.Run(fmt.Sprintf("%s:%d", ac.Host, ac.Port))
}
