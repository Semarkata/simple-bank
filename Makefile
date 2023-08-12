postgres:
	docker run --name postgre -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

destroy:
	docker rm -f postgre

start:
	docker start postgre

stop :
	docker stop postgre

createdb:
	docker exec -it postgre createdb --username=root --owner=root simple_bank
	docker exec -it postgre psql -U root -d simple_bank -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"

dropdb:
	docker exec -it postgre dropdb --username=root simple_bank

connect:
	docker exec -it postgre bash

migrateup:
	migrate -path DataBase/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path DataBase/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

.PHONY: postgres destroy start stop createdb dropdb connect migrateup migratedown