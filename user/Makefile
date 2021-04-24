machine_ip := $(shell ifconfig eth0 | grep 'inet' | cut -d: -f2 | awk '{print $$2}' | grep .)

build:
	docker build . -t ujinjinjin/user:latest
	docker image prune --filter label=stage=intermediate -f

push:
	docker push ujinjinjin/user:latest

publish:
	make build
	make push

run:
	docker run\
           -e SERVICE_ADDRESS=$(SERVICE_ADDRESS)\
           -e SERVICE_PORT=$(SERVICE_PORT)\
           -e DB_HOST=$(DB_HOST)\
           -e DB_USER=$(DB_USER)\
           -e DB_PASSWORD=$(DB_PASSWORD)\
           -e DB_NAME=$(DB_NAME)\
           -p 10000:$(SERVICE_PORT)\
           --add-host localhost:$(machine_ip)\
           ujinjinjin/user:latest

local_run:
	make build
	make run -e SERVICE_ADDRESS=""\
             -e SERVICE_PORT="10000"\
             -e DB_HOST="localhost"\
             -e DB_USER="local_user"\
             -e DB_PASSWORD="qwer1234"\
             -e DB_NAME="user_service"

go_build:
	go mod download
	go build -o bin/main .

go_run:
	./bin/main --address "localhost" \
               --port 10000\
               --db-host "localhost"\
               --db-user "local_user"\
               --db-password "qwer1234"\
               --db-name "user_service"

protoc_update:
	protoc --go_out=./interface --go_opt=paths=source_relative \
 		   --go-grpc_out=./interface --go-grpc_opt=paths=source_relative \
 		   --proto_path=../protocol_buffers user_service.proto
