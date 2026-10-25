.PHONY: ui-install ui-dev ui-build run build

ui-install:
	cd web && npm install

ui-dev:
	cd web && npm run dev

ui-build:
	cd web && npm run build

build: ui-build
	go build -o bin/lder .

run: ui-build
	go run .
