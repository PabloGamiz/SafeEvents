# Global about the project
VERSION=0.1.0
REPO=safe-events
PROJECT=backend
# Mysql variables
MYSQL_CONTAINER_NAME=safe-events-mysql
BACKEND_CONTAINER_NAME=safe-events

build:
	docker build -t ${REPO}/${PROJECT}:${VERSION} -f ./docker/safe-events/dockerfile .

run:
	go run ./cmd/safe-events/main.go

test:
	go clean -testcache
	go test -v ./...

deploy:
	docker-compose --env-file ./.env -f docker-compose.yaml up --remove-orphans -d

undeploy:
	docker-compose -f docker-compose.yaml down

mysql:
	docker logs ${MYSQL_CONTAINER_NAME} 2>&1 | grep GENERATED
	docker exec -it ${MYSQL_CONTAINER_NAME} mysql -uroot -p

dummy:
	docker logs ${BACKEND_CONTAINER_NAME} 2>&1 | grep ": Got a cookie"