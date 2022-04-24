deps:
	go mod download
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

generate: deps
	swagger generate client --target=./netbox --copyright-file=copyright_header.txt

clean:
	rm -rf netbox/client netbox/models

integration:
	go test ./... -tags=integration
