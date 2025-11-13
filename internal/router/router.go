package router

import (
	"go-fridge/internal/category"
	"go-fridge/internal/item"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	_ "go-fridge/internal/docs"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	itemRepo := item.NewRepository(db)
	itemService := item.NewService(itemRepo)
	itemHandler := item.NewHandler(itemService)

	categoryRepo := category.NewRepository(db)
	categoryService := category.NewService(categoryRepo)
	categoryHandler := category.NewHandler(categoryService)

	api := r.Group("/api")

	items := api.Group("/items")
	itemHandler.RegisterRoutes(items)

	categories := api.Group("/categories")
	categoryHandler.RegisterRoutes(categories)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
