package handlers

import (
	"awesomeProject/book/v3/config"
	"awesomeProject/book/v3/controllers"
	"awesomeProject/book/v3/models"
	"awesomeProject/book/v3/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 声明全局变量
var Book = &BookApiHandler{}

type BookApiHandler struct {
}

func (h *BookApiHandler) Registry(r gin.IRouter) {
	r.GET("/api/books", h.listBook)

	r.POST("/api/books", h.createBook)

	r.GET("/api/books/:bn", h.getBook)

	r.PUT("/api/books/:bn", h.updateBook)

	r.DELETE("/api/books/:bn", h.deleteBook)
}

func (h *BookApiHandler) listBook(ctx *gin.Context) {
	set := &models.BookSet{}

	pn, ps := 1, 20

	pageNumber := ctx.Query("page_number")
	if pageNumber != "" {
		pnInt, err := strconv.ParseInt(pageNumber, 10, 64)
		if err != nil {
			response.Failed(ctx, err)
			return
		}
		pn = int(pnInt)
	}

	pageSize := ctx.Query("page_size")
	if pageSize != "" {
		psInt, err := strconv.ParseInt(pageSize, 10, 64)
		if err != nil {
			response.Failed(ctx, err)
			return
		}
		ps = int(psInt)
	}

	query := config.DB().Model(&models.Book{})

	kws := ctx.Query("keywords")
	if kws != "" {
		query = query.Where("title LIKE ?", "%"+kws+"%")
	}

	offset := (pn - 1) * ps
	if err := query.Count(&set.Total).Offset(int(offset)).Limit(int(ps)).Find(&set.Items).Error; err != nil {
		response.Failed(ctx, err)
		return
	}
	response.OK(ctx, set)
}

func (h *BookApiHandler) createBook(ctx *gin.Context) {
	bookSpecInstance := &models.BookSpec{}

	if err := ctx.BindJSON(bookSpecInstance); err != nil {
		response.Failed(ctx, err)
		return
	}

	book, err := controllers.GetBookService().CreateBook(ctx.Request.Context(), bookSpecInstance)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	response.OK(ctx, book)
}

func (*BookApiHandler) getBook(ctx *gin.Context) {
	bnInt, err := strconv.ParseInt(ctx.Param("bn"), 10, 64)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	book, err := controllers.GetBookService().GetBook(ctx, controllers.NewGetBookRequest(int(bnInt)))
	if err != nil {
		response.Failed(ctx, err)
	}

	response.OK(ctx, book)

}

func (*BookApiHandler) updateBook(ctx *gin.Context) {
	bnStr := ctx.Param("bn")

	bn, err := strconv.ParseInt(bnStr, 10, 64)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	bookInstance := &models.Book{
		Id: uint(bn),
	}

	if err := ctx.BindJSON(bookInstance); err != nil {
		response.Failed(ctx, err)
		return
	}

	if err := config.DB().Where("id = ?", bookInstance.Id).Updates(bookInstance).Error; err != nil {
		response.Failed(ctx, err)
	}

	response.OK(ctx, bookInstance)
}

func (*BookApiHandler) deleteBook(ctx *gin.Context) {
	if err := config.DB().Where("id = ?", ctx.Param("bn")).Delete(&models.Book{}).Error; err != nil {
		response.Failed(ctx, err)
		return
	}

	response.OK(ctx, "ok")
}
