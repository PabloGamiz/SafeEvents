# Global about the project
VERSION=0.1.0
REPO=PabloGamiz
PROJECT=SafeEvents-Backend
# Volume varaibles
ROOT="${PWD}"
VOLUME_PATH=/tmp/safe-events
# Mysql variables
MYSQL_CONTAINER_NAME=mysql
MYSQL_VOLUME_PATH="${VOLUME_PATH}/mysql"
# PHPMyAdmin variables
MYADMIN_CONTAINER_NAME=myadmin
# DATABASE network
DB_NETWORK_NAME=${PROJECT}.network.db

build:
	docker build -t ${REPO}/${PROJECT}:${VERSION} -f ./docker/safe-events/dockerfile .

run:
	go run ./cmd/safe-events/main.go

test:
	go test ./...