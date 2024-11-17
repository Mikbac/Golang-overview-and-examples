package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSample(ctx *gin.Context) {
	testValue, _ := ctx.GetQuery("test") // query param
	fmt.Println(testValue)

	response := map[string]string{}
	response["test"] = "better " + testValue

	ctx.JSON(http.StatusOK, response)
}

func PostSample(ctx *gin.Context) {
	var message messageModel
	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response := map[string]string{}
	response["answerMessage"] = message.Message1

	ctx.JSON(http.StatusOK, response)
}
