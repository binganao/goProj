package main

import (
	context "context"
	"log"
	"main/product"
	"net"
	"time"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

type server struct {
	productMap map[string]*product.Product
}

func (s *server) AddProduct(ctx context.Context, req *product.Product) (resp *product.ProductId, err error) {
	resp = &product.ProductId{}
	out := uuid.NewV4()

	req.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*product.Product)
	}
	s.productMap[req.Id] = req
	resp.Value = req.Id
	return
}

func (s *server) GetProduct(ctx context.Context, req *product.ProductId) (resp *product.Product, err error) {
	if s.productMap == nil {
		s.productMap = make(map[string]*product.Product)
	}

	resp = s.productMap[req.Value]
	return
}

func srv() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	product.RegisterProductInfoServer(s, &server{})
	log.Println("Starting server on port 8080")
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}

const (
	address = "localhost:8080"
)

func cli() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := product.NewProductInfoClient(conn)
	ctx := context.Background()

	id := AddProduct(ctx, client)
	GetProduct(ctx, client, id)
}

func AddProduct(ctx context.Context, client product.ProductInfoClient) (id string) {
	Mac := &product.Product{Name: "MacBook 2019", Description: "From Apple"}
	productId, err := client.AddProduct(ctx, Mac)
	if err != nil {
		panic(err)
	}
	log.Println("Product added with id: ", productId.Value)
	return productId.Value
}

func GetProduct(ctx context.Context, client product.ProductInfoClient, id string) {
	product, err := client.GetProduct(ctx, &product.ProductId{Value: id})
	if err != nil {
		panic(err)
	}
	log.Println("Product: ", product)
}

func GRPCRun() {
	go srv()
	cli()

	<-time.After(time.Second * 3)
	log.Println("Done")
}
