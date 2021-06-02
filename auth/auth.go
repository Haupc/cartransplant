package main

import (
	"log"
	"net"

	"github.com/haupc/cartransplant/auth/controller"
	"github.com/haupc/cartransplant/auth/middleware"
	auth "github.com/haupc/cartransplant/auth/rpc"
	"github.com/haupc/cartransplant/cache"
	"github.com/haupc/cartransplant/grpcproto"
	"google.golang.org/grpc"

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
	go serveRPCServer()
	r := gin.Default()

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/logout", middleware.AuthorizeJWTFirebase(), authController.Logout)
		authRoutes.POST("/refesh-token", middleware.AuthorizeJWTFirebase(), authController.RefeshToken)
	}

	// resourceRoutes := r.Group("/geometry", middleware.Authorize())
	geometryRoutes := r.Group("/geometry")
	{
		// resourceRoutes.GET("/crawl-category", crawlController.CrawlCategory)
		geometryRoutes.GET("/current-address", middleware.AuthorizeJWTFirebase(), geometryController.GetCurrentAddress)
		geometryRoutes.GET("/get-route", geometryController.GetRouting)
		geometryRoutes.GET("/search-address", middleware.AuthorizeJWTFirebase(), geometryController.SearchAddress)

	}
	carRoutes := r.Group("/car", middleware.AuthorizeJWTFirebase())
	{
		carRoutes.POST("/register-trip", carController.RegisterTrip)
		carRoutes.POST("/find-trip", carController.FindTrip)
		carRoutes.POST("/register-car", carController.RegisterCar)
		carRoutes.PUT("/update-car", carController.UpdateCar)
		carRoutes.POST("/delete-car", carController.DeleteCar)
		carRoutes.GET("/list-my-car", carController.ListMyCar)
		carRoutes.GET("/user/list-trip", carController.ListUserTrip)
		carRoutes.POST("/take-trip", carController.TakeTrip)
	}
	r.Run(":8080")
}

func serveRPCServer() {
	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	grpcproto.RegisterAuthServer(grpcServer, auth.NewAuthServer())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
