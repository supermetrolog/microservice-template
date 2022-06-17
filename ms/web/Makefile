.PHONY: build
build:
	(if exist build rd -s -q build) && mkdir build && go build -o build/web.exe -v cmd/main.go

.PHONY: run
run:
	go run cmd/main.go

serve:
	build/generator.exe

.DEFAULT_GOAL := build