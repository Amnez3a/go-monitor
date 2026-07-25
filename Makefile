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
BIN_DIR=bin/
INSTALL_DIR=$(HOME)/.local/bin/
INSTALL_PATH=$(INSTALL_DIR)$(NAME)

build:
	go build -o $(BIN_DIR)$(NAME) main.go

build-docker:
	docker build -t $(NAME) .

run: build
	./$(BIN_DIR)$(NAME) --file servers.json

run-docker: build-docker
	docker run go-monitor --file servers.json

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
	rm -f bin/

docker-clean:
	docker rmi go-monitor

help:
	@echo "Доступные команды:"
	@echo "  make build        - собрать бинарник"
	@echo "  make build-docker - собрать в docker контейнер"
	@echo "  make run          - запустить бинарь"
	@echo "  make run-docker   - запустить docker образ"
	@echo "  make clean        - удалить bin/"
	@echo "  make install      - установить go-monitor в $(HOME)/.local/bin"
	@echo "  make uninstall    - удалить go-monitor из $(HOME)/.local/bin"
