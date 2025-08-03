include .env

build:
	go build -o bin/core
	chmod +x bin/core
	bin/core

air:
	air -c .air.toml

swag:
	swag init --parseDependency

mockery:
	mockery