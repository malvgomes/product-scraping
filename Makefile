up:
	docker-compose -f ./docker-compose.yml up --build

down:
	docker-compose down -v

mysql:
	docker-compose -f ./docker-compose.yml up --build mysql