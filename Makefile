.EXPORT_ALL_VARIABLES:
MONGO_URI:=mongodb://root:example@localhost:27017
DB_NAME:=evoke365-local
COLLECTION_NAME:=user

setup:
	brew tap go-swagger/go-swagger
	brew install go-swagger

gen:
	swagger generate server -f ./openapi/spec.yaml

run:
	go run ./cmd/evoke365-net-open-api-spec-server/main.go --host=localhost --port=8080 --tls-key /etc/openssl/ca.key --tls-certificate /etc/openssl/ca.crt