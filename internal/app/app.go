package app

import (
	"car-service/db"
	"car-service/internal/auth"
	"car-service/internal/broker"
	"car-service/internal/handler"
	"car-service/internal/model"
	"car-service/internal/repo"
	"car-service/internal/service"
	"log"
	"net"

	grpcapi "car-service/internal/transport/grpc"
	pb "car-service/internal/transport/grpc/vehiclepb"

	_ "car-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run() error {
	db.Connect()
	db.DB.AutoMigrate(&model.Vehicle{})
	db.DB.AutoMigrate(&model.User{})

	broker.InitWriter("vehicle.created")

	vehicRepo := repo.NewVehicleRepository(db.DB)
	vehicService := service.NewVehicleService(vehicRepo)

	vehicleHandler := handler.NewVehicleHandler(vehicService)

	authRepo := repo.NewAuthRepository(db.DB)
	authService := service.NewAuthService(authRepo)

	authHandler := auth.NewAuthHandler(authService)

	go func() {

		lis, err := net.Listen("tcp", ":9090")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		log.Println("Before grpc.Serve")

		grpcServer := grpc.NewServer()
		pb.RegisterVehicleServiceServer(grpcServer, grpcapi.NewVehicleGRPCServer(vehicService))

		reflection.Register(grpcServer)

		log.Println("Starting gRPC server")

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	r := gin.Default()
	handler.RegisterRoutes(r)
	vehicleHandler.Register(r)
	authHandler.Reg(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r.Run(":8080")
}
