package response

import (
	"awesomeProject/book/v3/exception"
	"github.com/gin-gonic/gin"
)

// 当前请求成功的时候，我们应用返回的数据
// 1. {code: 0, data: {}}
// 2. 正常直接返回数据, Restful接口 怎么知道这些请求是成功还是失败喃? 通过HTTP判断 2xx
// 如果后面 所有的返回数据 要进过特殊处理，都在这个函数内进行扩展，方便维护，比如 数据脱敏
func OK(ctx *gin.Context, data any) {
	ctx.JSON(200, data)
	ctx.Abort()
}

// 当前请求失败的时候, 我们返回的数据格式
// 1. {code: xxxx, data: null, message: "错误信息"}
// 请求HTTP Code 非 2xx 就返回我们自定义的异常
//
//	{
//		"code": 404,
//		"message": "book 1 not found"
//	}
func Failed(ctx *gin.Context, data any) {
	if e, ok := data.(*exception.ApiException); ok {
		ctx.JSON(e.HttpCode, e)
		ctx.Abort()
		return
	}
}
