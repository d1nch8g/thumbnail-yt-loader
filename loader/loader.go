package loader

import (
	"fmt"
	"io"
	"net/http"
)

func LoadThumbnail(uuid string) ([]byte, error) {
	link := fmt.Sprintf("https://img.youtube.com/vi/%s/1.jpg", uuid)
	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(resp.Body)
}
