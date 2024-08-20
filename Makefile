
.PHONY: gen
gen:
	mkdir -p gen
	protoc --proto_path=. --go_out=gen --go_opt=paths=source_relative service.proto
