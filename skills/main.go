package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"os"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
type BookSet struct {
	Total int64 `json:"total"`
	// ???Item声明什么意思
	Items []*Book `json:"items"`
}

type Book struct {
	Id uint `json:"id" gorm:"primaryKey;column;id"`

	// ??这是什么语法
	BookSepc
}

type BookSepc struct {
	Title  string  `json:"title" gorm:"column:title;type:varchar(200)" validate:"required"`
	Author string  `json:"author" gorm:"column:author;type:varchar(200);index" validate:"required"`
	Price  float64 `json:"price" gorm:"column:price;type:float" valiate:"required"`
	//nil 是零值， false
	IsSale bool `json:"is_sale" gorm:"column:is_sale"`
}

// ???这个table确定module的形式
func (b *Book) TableName() string {
	return "books"
}

func setUpDatabase() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306/go18?charset=utf8m4&parseTime=True&loc=Local"
	//??? 应该是Gin固定用法
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Book{})
	return db.Debug()
}

var db = setUpDatabase()

// ???这里是什么声明
var h = &BookApiHandler()

type BookApiHandler struct {
}

func (h *BookApiHandler) ListBook(ctx *gin.Context) {
	//??? bookHandler
	set = &BookSet{}

	pn, ps = 1, 20
	pageNumber := ctx.Query("page_number")
	if pageNumber != "" {
		pnInt, err := strconv.ParseInt(pageNumber, 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
			return
		}
		pn = int(pnInt)
	}

	pageSize := ctx.Query("page_size")
	if pageSize != "" {
		psInt, err = strconv.parseInt(pageSize, 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
			return
		}
		ps = int(psInt)
	}

	query := db.Module(&Book{})

	kws := ctx.Query("keywords")
	// 模糊查询
	if kws != "" {
		query = query.Where("title LIKE ?", "%"+kws+"%")
	}

	offset := (pn - 1) * ps
	if err := query.Count(&set, Total).offset(int(offset)).Limit(int(ps)).Find(&set, Items).Error; err != nil {
		// ???怎么两个500
		ctx.JSON(500, gin.H{"code": 500, "message": err.Error()})
		return
	}

	cts.JSON(200, set)
}

func (h *BookApiHandler) CteateBook(ctx *gin.Context) {

	// ???这里是干嘛
	bookSpecInstance := &BookSpec{}

	//???这个是获取request中的对象
	if err := ctx.BindJSON(bookSepcInstance); err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}

	// ????
	bookInstance := &Book(BookSpec: *bookSpecInstance)
	if err := db.Save(bookInstance).Error; err != nil {
		ctx.JSON(500, gin.H{"code":500, "message": err.Error()})
		return
	}
	//!!!返回结果
	ctx.JSON(http.StatusCreated, bookSpecInstance)
}

func (h *BookApiHandler) GetBook(ctx *gin.Context) {
	bookInstance := &Book{}
	if err := db.Where("id = ?", ctx.Param("bn").Take(bookInstance).Error; err != nil {
		ctx.JSON(400, gin.H{"code":400, "message": err.Error()})
		return
	}
	ctx.JSON(200, bookInstance)
}


func (h *BookApiHandler) UpdateBook(ctx *gin.Context) {
	bnStr := ctx.Param("bn")
	bn, err := strconv.ParseInt(bnStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}

	// ??? 声明id部分
	bookInstance := &Book{
		Id: uint(bn),
	}
	if err := ctx.BindJSON(&bookInstance.BookSpec); err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}

	if err := db.Where("id = ?", bookInstance.Id).Updates(bookInstance).Error; err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}
	ctx.JSON(200, bookInstance)
}

func (h *BookApiHandler) DeleteBook(ctx *gin.Context) {
	if err := db.Delete(&Book{}).Error; err != nil {
		ctx.JSON(400, gin.H{"code":400, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, "ok")
}

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	server := gin.Default()

	server.GET("/api/books", h.ListBook)

	server.POST("/api/books", h.CteateBook)

	server.GET("/api/books/:bn", h.GetBook)

	server.PUT("/api/books/:bn", h.UpdateBook)

	server.DELETE("/api/books/:bn", h.DeleteBook)


	if err := server.Run(":8080"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
