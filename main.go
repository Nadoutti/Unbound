package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register routes
	routes.RegisterAuthRoutes(r)

	r.Run() // listen and serve on

}
