package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/smhdhsn/restaurant-order/internal/config"
	"github.com/smhdhsn/restaurant-order/internal/db"
	"github.com/smhdhsn/restaurant-order/internal/repository/mysql"
	"github.com/smhdhsn/restaurant-order/internal/server"
	"github.com/smhdhsn/restaurant-order/internal/server/handler"
	"github.com/smhdhsn/restaurant-order/internal/server/resource"
	"github.com/smhdhsn/restaurant-order/internal/service"

	log "github.com/smhdhsn/restaurant-order/internal/logger"
	inventoryProto "github.com/smhdhsn/restaurant-order/internal/protos/edible/inventory"
	remoteRepository "github.com/smhdhsn/restaurant-order/internal/repository/remote"
)

// ctx holds application's context.
var ctx context.Context

// init will be called when this package is imported.
func init() {
	ctx = context.Background()
}

// main is the application's kernel.
func main() {
	// read configurations.
	conf, err := config.LoadConf()
	if err != nil {
		log.Fatal(err)
	}

	// create a database connection.
	dbConn, err := db.Connect(&conf.DB)
	if err != nil {
		log.Fatal(err)
	}

	// initialize auto migration.
	if err := mysql.InitMigrations(dbConn); err != nil {
		log.Fatal(err)
	}

	// make connection with external services.
	eConn, err := grpc.Dial(
		conf.Services[config.EdibleService].Address,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// instantiate gRPC clients.
	eiClient := inventoryProto.NewEdibleInventoryServiceClient(eConn)

	// instantiate repositories.
	eiRepo := remoteRepository.NewInventoryRepository(&ctx, eiClient)
	oRepo := mysql.NewOrderRepository(dbConn)

	// instantiate services.
	osServ := service.NewSubmissionService(eiRepo, oRepo)

	// instantiate handlers.
	osHand := handler.NewSubmitHandler(osServ)

	// instantiate resources.
	oRes := resource.NewOrderResource(osHand)

	// instantiate gRPC server.
	s, err := server.NewServer(&conf.Server, oRes)
	if err != nil {
		log.Fatal(err)
	}

	// listen and serve.
	if err := s.Listen(); err != nil {
		log.Fatal(err)
	}
}
