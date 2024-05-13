generate: generate-auth generate-version

generate-auth:
	rm -rf gen/go/auth
	mkdir -p gen/go/auth
	protoc \
		--proto_path=proto \
		--go_out=gen/go \
		--go_opt=paths=source_relative \
		--validate_out lang=go:gen/go \
		--validate_opt=paths=source_relative \
		proto/auth/*.proto


generate-version:
	rm -rf gen/go/version
	mkdir -p gen/go/version
	protoc \
		--proto_path=proto \
		--go_out=gen/go \
		--go_opt=paths=source_relative \
		--validate_out lang=go:gen/go \
		--validate_opt=paths=source_relative \
		proto/version/*.proto