package main

import "github.com/gin-gonic/gin"

func main(){
	engin := gin.Default()
	route(engin)
	engin.RunTLS(":443","./key/nudao.crt","./key/nudao.key")
}