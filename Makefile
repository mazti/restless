.PHONY: install
install:
	go build -o ./restless ./cmd
	mv ./restless $(GOPATH)/bin/restless

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down