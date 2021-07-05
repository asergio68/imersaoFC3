package main

import (
	"log"
	"net"

	"github.com/asergio68/imersaoFC3/goapp-grpc/entity"
	"github.com/asergio68/imersaoFC3/goapp-grpc/grpc/pb"
	"github.com/asergio68/imersaoFC3/goapp-grpc/grpc/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var ProductList = entity.NewProducts()

func main() {
	lis, err := net.Listen("tcp","localhost:50051")
	if err != nil {
		log.Print("Erro ao se conectar");
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	productService := service.NewProductGrpcService()
	productService.Products = ProductList
	pb.RegisterProductServiceServer(grpcServer, productService)
	grpcServer.Serve(lis)
}