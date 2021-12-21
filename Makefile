up:
	docker-compose -f ./docker-compose.yml up --build

down:
	docker-compose down -v

db:
	docker-compose -f ./docker-compose.yml up --build db