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
           ujinjinjin/user:latest

local_run:
	make run -e SERVICE_ADDRESS=""\
             -e SERVICE_PORT="10000"\
             -e DB_HOST="localhost"\
             -e DB_USER="local_user"\
             -e DB_PASSWORD="qwer1234"\
             -e DB_NAME="local_db"

go_build:
	go mod download
	go build -o bin/main .

go_run:
	./bin/main --address "localhost" \
               --port 10000\
               --db-host "localhost"\
               --db-user "local_user"\
               --db-password "qwer1234"\
               --db-name "local_db"