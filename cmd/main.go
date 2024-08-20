package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"thumbnail-yt-loader/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Video link example: mbeT8mpmtHA

func main() {
	if len(os.Args) > 1 {
		os.Args = os.Args[1:]
		for i := range os.Args {
			if strings.Contains(os.Args[i], "watch?v=") {
				os.Args[i] = strings.Split(os.Args[i], "watch?v=")[1]
			}
		}
		fmt.Println("Loading thumbnails for following uuid's:", os.Args)

		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}

		conn, err := grpc.NewClient("localhost:8092", opts...)
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		client := gen.NewThumbnailLoaderClient(conn)

		for _, uuid := range os.Args {
			resp, err := client.Load(context.Background(), &gen.ThumbnailRequest{
				Uuid: uuid,
			})
			if err != nil {
				panic(err)
			}
			os.WriteFile(fmt.Sprintf("%s.png", uuid), resp.Thumbnail, os.ModePerm)
		}
	}
}
