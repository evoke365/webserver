gen:
	swagger generate server -f ./openapi/spec.yaml

run:
	go run ./cmd/evoke365-net-open-api-spec-server/main.go --host=localhost --port=8080