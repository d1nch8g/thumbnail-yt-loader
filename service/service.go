package service

import (
	"context"
	"fmt"
	"thumbnail-yt-loader/cache"
	"thumbnail-yt-loader/gen"
	"thumbnail-yt-loader/loader"
)

type LoaderService struct {
	gen.ThumbnailLoaderServer
}

func (s *LoaderService) Load(ctx context.Context, r *gen.ThumbnailRequest) (*gen.ThumbnailResponse, error) {
	result, err := cache.CheckThumbnail(r.Uuid)
	if err != nil {
		result, err = loader.LoadThumbnail(r.Uuid)
		if err != nil {
			return nil, err
		}
		err = cache.StoreThumbnail(r.Uuid, result)
		if err != nil {
			fmt.Printf("unable to store cache for uuid: %v", err)
		}
	}
	fmt.Printf("Thumbnail for %s loaded...\n", r.Uuid)
	return &gen.ThumbnailResponse{Thumbnail: result}, nil
}
