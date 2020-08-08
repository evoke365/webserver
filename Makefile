.EXPORT_ALL_VARIABLES:
GO111MODULE:=on
MONGO_URI:=mongodb://root:example@localhost:27017
DB_NAME:=evoke365-local
COLLECTION_NAME:=user

setup:
	brew tap go-swagger/go-swagger
	brew install go-swagger

gen:
	swagger generate server -f ./openapi/spec.yaml

run:
	go run ./cmd/evoke365-net-open-api-spec-server/main.go --tls-port=8080 --tls-key openssl/ca.key --tls-certificate openssl/ca.crt

test:
	go test ./...