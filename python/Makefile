start: build run

build:
	cd app && docker build -t madlib-python-service .

run:
	cd app && docker run -d -p 8000:8000 madlib-python-service

remove: remove-container remove-image

remove-container:
	docker ps -a | grep "madlib-python-service" | awk '{print $$1}' | xargs docker rm -f

remove-image:
	docker images -a | grep "madlib-python-service" | awk '{print $$3}' | xargs docker rmi