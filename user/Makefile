build_go:
	go mod download
	go build -o bin/main .

run_go:
	./bin/main

build:
	docker build . -t ujinjinjin/user:latest
	docker image prune --filter label=stage=intermediate -f

push:
	docker push ujinjinjin/user:latest

publish:
	make build
	make push

run:
	docker run -p 10000:10000 ujinjinjin/user:latest