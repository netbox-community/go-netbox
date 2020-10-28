swagger=docker run --rm -it -e GOPATH=$(HOME)/go:/go -v $(HOME):$(HOME) -w $$(pwd) quay.io/goswagger/swagger:v0.25.0
generate:
	$(swagger) generate client --target=./netbox --spec=./swagger.processed.json --copyright-file=./copyright_header.txt

preprocess:
	bash preprocess.sh swagger.json > swagger.processed.json

clean:
	rm -rf netbox/client netbox/models

integration:
	go test ./... -tags=integration
