create-migration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate-up: create-db
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up

migrate-down: create-db
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down

create-db:
	docker compose exec mysql mysql -uroot -proot -e "CREATE DATABASE IF NOT EXISTS courses;"

.PHONY: create-migration migrate-up migrate-down create-db