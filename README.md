# Youtube thumbnail loader service

Simple golang service with gRPC API to load youtube thumbnails. With applied command line utility used to load thumbnails. Server provided in `docker`.

## Run locally

### Run service

To build and run service locally you have to have `golang` installed. Run following command:


```shell
go run .
```

### Run client

To run file located in `cmd` folder, 

```
go run cmd/main.go mbeT8mpmtHA hu-q2zYwEYs
```

### Run service in docker

Execute following command:

```
docker build -t ion.lc/d1nch8g/thumbnail-yt-loader .
docker run -p 8092:8092 ion.lc/d1nch8g/thumbnail-yt-loader
```
