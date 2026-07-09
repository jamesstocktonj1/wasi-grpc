generate:
	@echo "Generating gRPC bindings..."
	@mkdir -p gen
	@protoc \
		--go_out=. --go_opt=paths=import \
		--go-grpc_out=. --go-grpc_opt=paths=import \
		proto/greet.proto

clean:
	@rm -rf gen