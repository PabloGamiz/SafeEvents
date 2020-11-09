# Global about the project
VERSION=0.1.0
REPO=safe-events
PROJECT=backend
# Mysql variables
MYSQL_CONTAINER_NAME=safe-events-mysql
build:
	docker build -t ${REPO}/${PROJECT}:${VERSION} -f ./docker/safe-events/dockerfile .
run:
	go run ./cmd/safe-events/main.go
test:
	go clean -testcache
	go test -v ./...

deploy:
	docker-compose -f docker-compose.yaml up
	docker-compose --env-file ./.env -f docker-compose.yaml up --remove-orphans

undeploy:
	docker-compose -f docker-compose.yaml down