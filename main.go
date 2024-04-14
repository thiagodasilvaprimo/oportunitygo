package main

import "github.com/gin-gonic/gin"

func main() {
	// Inicializa o Router utilizando as configuracoes default do gin
	r := gin.Default()

	// DEFININDO UMA ROTA
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}