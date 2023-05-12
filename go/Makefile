start: build run

build:
	cd app && docker build -t madlib-go-service .

run:
	cd app && docker run -d -p 8080:8080 madlib-go-service

remove: remove-container remove-image

remove-container:
	docker ps -a | grep "madlib-go-service" | awk '{print $$1}' | xargs docker rm -f

remove-image:
	docker images -a | grep "madlib-go-service" | awk '{print $$3}' | xargs docker rmi