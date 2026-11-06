package main

import (
	"go-api/application/useCases"
	"go-api/infra/db"
	"go-api/infra/repositories"
	"go-api/presentation"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	productRepository := repository.NewProductRepository(dbConnection)
	productUseCase := useCases.NewProductUseCase(productRepository)
	productController := controller.NewProductController(productUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.CreateProduct)
	server.GET("/product/:productId", productController.GetProductById)

	server.Run(":8000")
}
