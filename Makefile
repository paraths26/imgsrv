NAME=imgsrv
VERSION=`git describe --abbrev=0 --tag`

all: build run


build:
	docker build -t $(NAME):ocr .

run:
	docker-compose up -d