package main

import (
	"fmt"
	"log"
	"net"
	"thumbnail-yt-loader/gen"
	"thumbnail-yt-loader/service"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8092))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	fmt.Println("Starting new thumbnail loader service...")

	grpcServer := grpc.NewServer(opts...)

	gen.RegisterThumbnailLoaderServer(grpcServer, &service.LoaderService{})
	grpcServer.Serve(lis)
}
