up:
	docker-compose -f ./docker-compose.yml up --build

down:
	docker-compose down -v

db:
	docker exec -ti product-scraping_db_1 mysql -u root -p1234

test:
	docker exec -ti product-scraping_app_1  go test -cover ./...

