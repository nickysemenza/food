VERSION          := $(shell git describe --tags --always --dirty="-dev")
DATE             := $(shell date '+%Y-%m-%d-%H%M UTC')
VERSION_FLAGS    := -ldflags='-X "main.Version=$(VERSION)" -X "main.BuildTime=$(DATE)"'

GCP_PROJECT := cloudrun1-278204
GCR_IMAGE := gcr.io/$(GCP_PROJECT)/food-backend:$(VERSION)

deploy:
	gcloud builds submit --tag $(GCR_IMAGE)
	terraform apply -var-file="prod.tfvars" -var="image_name=$(GCR_IMAGE)"

dev: bin/food
	./bin/food

.PHONY: all
all: bin/food

bin/%: $(shell find . -type f -name '*.go' | grep -v '_test.go')
	@mkdir -p $(dir $@)
	go build $(VERSION_FLAGS) -o $@ ./cmd/$(@F)

test: unit-test lint

dev-env:
	docker-compose up -d db jaeger

bin/revive:
	@mkdir -p $(dir $@)
	go build -o $@ ./vendor/github.com/mgechev/revive
bin/migrate:
	@mkdir -p $(dir $@)
	go build -tags 'postgres' -o $@ ./vendor/github.com/golang-migrate/migrate/v4/cmd/migrate
unit-test: 
	go test -v -race -cover ./...
lint: bin/revive
	bin/revive -config revive.toml -formatter=friendly -exclude=vendor/... ./... || (echo "lint failed"; exit 1)	

IMAGE := nicky/food-backend
docker-build:
	docker build -t $(IMAGE) .
docker-dev: docker-build

docker-push: docker-build
	docker push $(IMAGE):latest

dev-db:
	pgcli postgres://food:food@localhost:5555/food
new-migrate/%: bin/migrate
	mkdir -p migrations
	./bin/migrate create -dir migrations -ext sql $(@F)
migrate: bin/migrate
	./bin/migrate -source file://migrations -database postgres://food:food@localhost:5555/food?sslmode=disable up
migrate-down: bin/migrate
	./bin/migrate -source file://migrations -database postgres://food:food@localhost:5555/food?sslmode=disable down

.PHONY: generate-graphql-go
generate-graphql-go: 
	go generate ./graph
generate-graphql-ts:
	cd ui && yarn run generate
	cd cli && yarn run generate

graphql: generate-graphql-go generate-graphql-ts