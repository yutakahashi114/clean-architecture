package rest

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/yutakahashi114/clean-architecture/controller/rest/handler"
	rFirestore "github.com/yutakahashi114/clean-architecture/infrastructure/firestore"
	rGrpc "github.com/yutakahashi114/clean-architecture/infrastructure/grpc"
	"github.com/yutakahashi114/clean-architecture/infrastructure/postgres"
	"github.com/yutakahashi114/clean-architecture/usecase"
	"google.golang.org/grpc"
)

var port = os.Getenv("PORT")

// Serve .
func Serve() {

	router := gin.Default()

	// postgres
	db, err := postgres.NewDB(
		postgres.DBConfig{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Name:     os.Getenv("DB_NAME"),
		},
	)
	if err != nil {
		log.Printf("failed to connect postgres: %v", err)
		return
	}
	defer db.Close()
	restaurantRepository := postgres.NewRestaurantRepository(db)

	// firestore
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Printf("failed to connect firestore: %v", err)
		return
	}
	defer client.Close()
	_ = rFirestore.NewRestaurantRepository(client)

	// grpc
	conn, err := grpc.Dial(
		os.Getenv("GRPC_HOST"),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Printf("failed to connect grpc: %v", err)
		return
	}
	defer conn.Close()
	_ = rGrpc.NewRestaurantRepository(conn)

	restaurantUseCase := usecase.NewRestaurantUseCase(restaurantRepository)
	restaurantHandler := handler.NewRestaurantHandler(restaurantUseCase)

	router.GET("/restaurant/:id", restaurantHandler.GetByID)

	log.Println("REST server start: port " + port)
	if err := router.Run(":" + port); err != nil {
		log.Printf("failed to serve: %v", err)
		return
	}
}
