package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.New()
	err := server.Run()
	if err != nil {
		return
	}
}
