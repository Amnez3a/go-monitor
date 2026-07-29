go-monitor: go
	if command -v apt &> /dev/null; then
		echo "system use apt"
		sudo apt update && sudo apt install go
	elif command -v dnf &> /dev/null; then
		echo "system use dnf"
		sudo dnf clean metadata && sudo dnf makecache && sudo dnf install go
	elif command -v pacman &> /dev/null; then
		echo "system use pacman"
		sudo pacman -Sy go
	elif command -v zypper &> /dev/null; then
		echo "system use zypper"
		sudo zypper refresh && zypper install go
	fi

.PHONY: build run clean install uninstall

NAME=go-monitor
BIN_DIR=bin
INSTALL_DIR=$(HOME)/.local/bin/
INSTALL_PATH=$(INSTALL_DIR)$(NAME)

build:
	go build -o $(BIN_DIR)/$(NAME) main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/$(NAME) main.go

build-macos:
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/$(NAME) main.go

build-docker:
	docker build -t $(NAME) .

build-i386:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o bin/go-monitor .

run: build
	./$(BIN_DIR)/$(NAME)

run-windows: build-windows
	.\$(BIN_DIR)\$(NAME) # powershell

run-macos: build-macos
	./$(BIN_DIR)/$(NAME)

run-docker: 
	docker run -ti --rm -v ./servers.json:/app/servers.json $(NAME)
	
install:
	@mkdir -p $(INSTALL_DIR)
	go build -o $(INSTALL_PATH) main.go
	@echo "installed to $(INSTALL_PATH)"
	@echo "set path $(HOME)/.local/bin"

uninstall:
	@if [ -f $(INSTALL_PATH) ]; then \
		rm -f $(INSTALL_PATH); \
		echo "uninstalled from $(INSTALL_DIR)/"; \
	else \
		echo "nothing to uninstall (file not found)"; \
	fi

clean:
	rm -f $($BIN_DIR)/*

docker-clean:
	docker rmi -f $(NAME)

help:
	@echo "Commands:"
	@echo "  make build        - build binary"
	@echo "  make build-docker - make docker image"
	@echo "  make run          - run bin/go-monitor"
	@echo "  make run-docker   - run docker image"
	@echo "  make clean        - remove bin/"
	@echo "  make install      - install go-monitor to $(HOME)/.local/bin"
	@echo "  make uninstall    - uninstall go-monitor from $(HOME)/.local/bin"