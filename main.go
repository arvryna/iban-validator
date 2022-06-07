package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func launchServer() {
	server := gin.New()
	server.Use(gin.Recovery()) // default Middlewhere that catches panic and returns 500 as default
	server.POST("/validators/iban", func(ctx *gin.Context) {
		ctx.JSON(400, gin.H{
			"message": "Invalid Iban number",
		})
	})

	server.Run(":9090")
}

func main() {
	fmt.Println("IBAN-Validator")
	launchServer()
}
