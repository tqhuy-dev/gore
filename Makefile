grpc-dependency:
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
generate:
	protoc \
		-I ./proto \
		-I ./grpc-party/googleapis \
		-I ./grpc-party/ \
		--go_out ./pb --go_opt=paths=source_relative \
		--go-grpc_out ./pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out ./pb --grpc-gateway_opt=paths=source_relative \
		--validate_out=lang=go,paths=import:./pb \
		--openapiv2_out ./swagger --openapiv2_opt=logtostderr=true \
		proto/*.proto

generate-jaeger-proto:
	protoc \
		   -I ./proto \
		   -I ./grpc-party/ \
		   --go_out=./ --go_opt=paths=source_relative \
           --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
           proto/model.proto \
