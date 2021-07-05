package service

import (
	"context"

	"github.com/asergio68/imersaoFC3/goapp-grpc/entity"
	"github.com/asergio68/imersaoFC3/goapp-grpc/grpc/pb"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	Products *entity.Products
}

func NewProductGrpcService() (*ProductService){
	return &ProductService{}
}

func (p *ProductService) CreateProduct(ctx context.Context, in *pb.Product) (*pb.ProductResult, error) {
	pr := entity.NewProduct(in.GetName())
	p.Products.Add(pr)

	return &pb.ProductResult{
		Id: pr.ID,
		Name: pr.Name,
	}, nil
}