deps:
	go mod download
	GO111MODULE=off go get -u github.com/myitcv/gobin

generate: deps
	gobin -m -run github.com/go-swagger/go-swagger/cmd/swagger generate client --copyright-file=copyright_header.txt

clean:
	rm -rf netbox/client netbox/models

integration:
	go test ./... -tags=integration
