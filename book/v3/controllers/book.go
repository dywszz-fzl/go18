package controllers

import (
	"awesomeProject/book/v3/config"
	"awesomeProject/book/v3/exception"
	"awesomeProject/book/v3/models"
	"awesomeProject/skills/ioc"
	"context"

	"gorm.io/gorm"
)

func GetBookService() *BookController {
	return ioc.Controller.Get("book_controller").(*BookController)
}

func init() {
	ioc.Controller.Registry("book_controller", &BookController{})
}

type BookController struct {
	ioc.ObjectImpl
}

func NewGetBookRequest(bookNumber int) *GetBookRequest {
	return &GetBookRequest{
		BookNumber: bookNumber,
		// RequestId string
	}
}

type GetBookRequest struct {
	BookNumber int
}

// 核心功能
// ctx: Trace, 支持请求的取消, request_id
// GetBookRequest 为什么要把他封装为1个对象, GetBook(ctx context.Context, BookNumber string), 保证你的接口的签名的兼容性
// BookController.GetBook(, "")
func (c *BookController) GetBook(ctx context.Context, in *GetBookRequest) (*models.Book, error) {
	config.L().Debug().Msgf("get book %d", in.BookNumber)

	bookInstance := &models.Book{}

	if err := config.DB().Where("id = ?", in.BookNumber).Take(bookInstance).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.ErrNotFound("Book number: %d not Found", in.BookNumber)
		}
		return nil, err
	}

	return bookInstance, nil
}

func (c *BookController) CreateBook(ctx context.Context, in *models.BookSpec) (*models.Book, error) {
	// 有没有能够检查某个字段是否是必须填
	// Gin 集成 validator这个库, 通过 struct tag validate 来表示这个字段是否允许为空
	// validate:"required"
	// 在数据Bind的时候，这个逻辑会自动运行
	// if bookSpecInstance.Author == "" {
	// 	ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
	// 	return
	// }

	bookInstance := &models.Book{BookSpec: *in}

	if err := config.DB().Save(bookInstance).Error; err != nil {
		return nil, err
	}

	return bookInstance, nil

}

func (c *BookController) UpdateBook() {
	// update(obj)
	// config.DB().Updates()
}

func (c *BookController) update(ctx context.Context, obj models.Book) error {
	// obj.UpdateTime = now()
	// config.DB().Updates()
	return nil
}
