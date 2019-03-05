package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func index(ctx *gin.Context){
	ctx.HTML(http.StatusOK,"./view/index.html",nil)
}
