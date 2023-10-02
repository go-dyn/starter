COMMAND=docker compose -f docker-compose.yml
APP_CONTAINER=app

.PHONY: up
up:
	$(COMMAND) up -d

.PHONY: down
down:
	$(COMMAND) down

.PHONY: bash
bash:
	$(COMMAND) exec $(APP_CONTAINER) sh

.PHONY: artisan
artisan: 
	$(COMMAND) exec $(APP_CONTAINER) go run . artisan $(args)

.PHONY: test
test:
	$(COMMAND) exec $(APP_CONTAINER) go test $(args)

.PHONY: go-get
go-get: 
	$(COMMAND) exec $(APP_CONTAINER) go get $(args)

.PHONY: packages-update
packages-update:
	$(COMMAND) exec $(APP_CONTAINER) go get -u ./...

.PHONY: build
build:
	go build -a --ldflags "-w -s" -gcflags=all="-l -B" -o hero .