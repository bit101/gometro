default:
	tmux_send "make build"

build:
	@go build main.go
	@./main

