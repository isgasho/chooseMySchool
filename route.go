package main

import "github.com/gin-gonic/gin"

func route(g *gin.Engine){
g.GET("/",index)
}
