package utils

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  msg,
		"data": data,
	})
}

func ErrorResponse(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
		"data": data,
	})
}

func UrlJoin(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return " "
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return " "
	}

	return baseUrl.ResolveReference(uri).String()
}
