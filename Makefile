.PHONY: all
all: build fmt vet lint test

APP=contacts
DB_USER=postgres
GLIDE_NOVENDOR=$(shell glide novendor)
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
DB_NAME=contacts_db
DB_USER=dev

APP_EXECUTABLE="./out/$(APP)"
APP_MIGRATIONS="./out/migrations"

setup:
	go get -u github.com/golang/lint/golint

clean:
	rm -f $(APP_EXECUTABLE)

build-deps:
	glide cc; glide install;

update-deps:
	glide update

compile:
	mkdir -p out/migrations
	cp -r ./migrations $(APP_MIGRATIONS)
	go build -o $(APP_EXECUTABLE)

build:	build-deps compile fmt vet

create_image:	clean setup build

install:
	go install ./...

fmt:
	go fmt $(GLIDE_NOVENDOR)

vet:
	go vet $(GLIDE_NOVENDOR)

db.setup:	db.create_user db.create db.migrate

db.create_user:
	createuser -s $(DB_USER) || true

db.create:
	createdb -O$(DB_USER) -Eutf8 $(DB_NAME)

db.migrate:
	$(APP_EXECUTABLE) migrate:run

db.rollback:
	$(APP_EXECUTABLE) migrate:rollback

db.drop:
	dropdb --if-exists -U$(DB_USER) $(DB_NAME)

db.reset:	db.drop db.create db.migrate

copy-config:
	cp application.yml.sample application.yml
