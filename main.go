package main

import (
	"net/http"

	"github.com/arvryna/iban-validator/internal/ibanparser"
	"github.com/gin-gonic/gin"
)

type IbanRequestBody struct {
	Iban string `json:"iban"`
}

func main() {
	server := gin.New()
	server.Use(gin.Recovery()) // default Middlewhere that catches panic and returns 500 as default
	server.POST("/validators/iban", func(ctx *gin.Context) {
		var ibanReqBody IbanRequestBody
		err := ctx.ShouldBindJSON(&ibanReqBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		parser := ibanparser.Init(ibanReqBody.Iban)
		if err := parser.Validate(); err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(200, gin.H{
				"message": "Valid IBAN",
			})
		}
	})
	server.Run(":9090")
}
