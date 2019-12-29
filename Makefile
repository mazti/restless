.PHONY: install
install:
	go install $(pwd)

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down