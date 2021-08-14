.EXPORT_ALL_VARIABLES:
GO111MODULE:=on
MONGO_URI:=mongodb://root:example@localhost:27017
DB_NAME:=evoke365-local
COLLECTION_NAME:=user

test:
	go test ./...

setup:
	brew tap go-swagger/go-swagger
	brew install go-swagger

key-gen:
	openssl genrsa -out openssl/key.pem
	openssl req -new -x509 -key openssl/key.pem -out openssl/ca.crt

gen:
	swagger generate server -f ./openapi/spec.yaml
	
run:
	go run ./cmd/evoke365-net-open-api-spec-server/main.go --tls-port=8080 --tls-key openssl/key.pem --tls-certificate openssl/ca.crt

local-deps:
	TMPDIR=/private$$TMPDIR docker-compose up