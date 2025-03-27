package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func LoggingMiddleware(ctx *gin.Context) {
	log.Printf("Request - Method: %s, Path: %s", ctx.Request.Method, ctx.Request.URL.Path)
	ctx.Next()
}

func AuthMiddleware() {

}
