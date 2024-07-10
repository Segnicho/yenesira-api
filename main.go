package main

import (
	"fmt"
	"yenesira-api/routes"
	"github.com/gin-gonic/gin"

)

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	fmt.Println("Starting server on port 8080")
	r.Run(":8080")
	
}