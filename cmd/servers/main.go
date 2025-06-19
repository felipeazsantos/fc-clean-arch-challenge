package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/configs"
	"github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/internal/infra/database"
	"github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/internal/infra/graph"
	"github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/internal/infra/grpc/pb"
	"github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/internal/infra/grpc/service"
	"github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/internal/infra/web"
	"github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/internal/infra/web/webserver"
	"github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := configs.LoadConfig("./cmd/servers/")
	if err != nil {
		log.Fatal("failed to load config: " + err.Error())
	}

	db, err := configs.ConnectDB(cfg)
	if err != nil {
		log.Fatal("cannot connect with DB:", err)
	}

	orderRepository := database.NewOrderRepository(db)
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository)
	listOrdersUseCase := usecase.NewListOrdersUseCase(orderRepository)

	webserver := webserver.NewWebServer(cfg.WebServerPort)
	webOrderHandler := web.NewWebOrderHandler(orderRepository)
	webserver.AddHandler("POST", "/orders/create", webOrderHandler.Create)
	webserver.AddHandler("GET", "/orders/list", webOrderHandler.ListOrders)

	fmt.Println("Starting web server on port:", cfg.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	orderService := service.NewOrderService(createOrderUseCase, listOrdersUseCase)
	pb.RegisterOrderServiceServer(grpcServer, orderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port:", cfg.GrpcServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GrpcServerPort))
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to listen on port %s:", cfg.GrpcServerPort), err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			CreateOrderUseCase: createOrderUseCase,
			ListOrdersUseCase:  listOrdersUseCase,
		},
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port:", cfg.GraphqlServerPort)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.GraphqlServerPort), nil)
}
