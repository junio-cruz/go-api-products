package main

import (
	"go-api/presentation/controller/ProductController"
	"go-api/db"
	"go-api/infra/repositories/ProductRepository"
	"go-api/application/useCases/ProductUseCases"
	"github.com/gin-gonic/gin"
)
	
func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)

	server.Run(":8000")
}