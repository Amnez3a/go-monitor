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

.PHONY: build run clean

build:
	go build -o bin/go-monitor .

build-docker:
	docker build -t go-monitor .

build-i386:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o bin/go-monitor .

run: build
	./bin/go-monitor

run: build-docker
	docker run -it go-monitor

run: build-i386
	./bin/go-monitor

clean:
	rm -rf bin/

docker-clean:
	docker rmi go-monitor

help:
	@echo "Доступные команды:"
	@echo "  make build  - собрать бинарник"
	@echo "  make build-docker - собрать в docker контейнер"
	@echo "  make build-i386 - собрать под i386"
	@echo "  make run    - запустить"
	@echo "  make clean  - удалить bin/"
