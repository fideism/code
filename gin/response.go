package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Make 返回Json格式 200
func Make(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

// MakePage 返回分页信息 200
func MakePage(c *gin.Context, data interface{}, page PageResponse) {
	Make(c, map[string]interface{}{
		"list":       data,
		"pagination": page,
	})
}

// Ok 正常
func Ok(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": message,
	})
}

// Created 201
func Created(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "保存记录成功",
	})
}

// CreatedWithData  201 并且返回data信息
func CreatedWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"code": http.StatusCreated,
		"data": data,
	})
}

// NoContent 204 not content
func NoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"code":    http.StatusNoContent,
		"message": "Not Found",
	})
}

// BadRequest 客户端错误
func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"message": message,
	})
}

// Unauthorized 请登录后访问
func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    http.StatusUnauthorized,
		"message": message,
	})
}

// Unprocessable 参数格式错误
func Unprocessable(c *gin.Context, message string) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"code":    http.StatusUnprocessableEntity,
		"message": message,
	})
}

// ServerError 服务端错误
func ServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusInternalServerError,
		"message": message,
	})
}

// NotFound 404 not found
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    http.StatusNotFound,
		"message": "Not Found",
	})
}

// Error 自定义错误
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

// CodeWithData 自定义错误
func CodeWithData(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
	})
}
