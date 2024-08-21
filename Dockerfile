FROM golang:1.22
WORKDIR /src
COPY . .
RUN go build -o /bin/server ./main.go

FROM alpine:3.20
COPY --from=0 /bin/server /bin/server
CMD ["/bin/server"]