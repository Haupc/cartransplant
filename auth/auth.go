package main

import (
	"github.com/haupc/cartransplant/auth/controller"
	"github.com/haupc/cartransplant/auth/middleware"
	"github.com/haupc/cartransplant/cache"

	"github.com/gin-gonic/gin"
)

var (
	authController = controller.GetAuthController()
	userCache      = cache.GetUserCache()
)

func main() {
	r := gin.Default()

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/logout", middleware.Authorize(), authController.Logout)
		authRoutes.POST("/refesh-token", middleware.Authorize(), authController.RefeshToken)
	}

	resourceRoutes := r.Group("/geometry", middleware.Authorize())
	{
		// resourceRoutes.GET("/crawl-category", crawlController.CrawlCategory)
		resourceRoutes.GET("/current-address")

	}
	r.Run()
}
