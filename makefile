.PHONY: all compile_backend compile_frontend build_docker test

openapi-generate:
	npx @openapitools/openapi-generator-cli generate -i api/v1/swagger.yaml -g typescript-axios -o src/portal/src/common
	npx @openapitools/openapi-generator-cli generate -i api/v1/swagger.yaml -g go-server -o src/common
