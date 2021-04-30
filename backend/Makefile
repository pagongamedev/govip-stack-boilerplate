ifeq ($(OS),Windows_NT)
SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command
endif

ifeq ($(OS),Windows_NT)
build:
	$$env:GO111MODULE="on"; $$env:API_ENV="localhost";
	go build .\main.go; .\main.go

run:
	$$env:GO111MODULE="on"; $$env:API_ENV="localhost";
	go run .\main.go
else

build:
	env GO111MODULE=on API_ENV=localhost \
	go build main.go

run:
	env GO111MODULE=on API_ENV=localhost \
	go run main.go

endif