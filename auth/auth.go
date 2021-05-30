package main

import (
	"github.com/haupc/cartransplant/auth/controller"
	"github.com/haupc/cartransplant/auth/middleware"
	"github.com/haupc/cartransplant/cache"

	"github.com/gin-gonic/gin"
)

var (
	authController     = controller.GetAuthController()
	geometryController = controller.GetGeometryController()
	carController      = controller.GetCarController()
	userCache          = cache.GetUserCache()
)

func main() {
	// config.GetDbConnection().AutoMigrate(&model.User{}, model.Role{}, model.Permission{})
	r := gin.Default()

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/logout", middleware.AuthorizeJWTFirebase(), authController.Logout)
		authRoutes.POST("/refesh-token", middleware.AuthorizeJWTFirebase(), authController.RefeshToken)
	}

	// resourceRoutes := r.Group("/geometry", middleware.Authorize())
	geometryRoutes := r.Group("/geometry", middleware.AuthorizeJWTFirebase())
	{
		// resourceRoutes.GET("/crawl-category", crawlController.CrawlCategory)
		geometryRoutes.GET("/current-address", geometryController.GetCurrentAddress)
		geometryRoutes.GET("/get-route", geometryController.GetRouting)
		geometryRoutes.GET("/search-address", geometryController.SearchAddress)

	}
	carRoutes := r.Group("/car", middleware.AuthorizeJWTFirebase())
	{
		carRoutes.POST("/register-trip", carController.RegisterTrip)
		carRoutes.POST("/find-trip", carController.FindTrip)
		carRoutes.POST("/register-car", carController.RegisterCar)
		carRoutes.PUT("/update-car", carController.UpdateCar)
		carRoutes.DELETE("/delete-car", carController.DeleteCar)
		carRoutes.GET("/list-my-car", carController.ListMyCar)
	}
	r.Run(":8080")
}
