
databaseName = science
databaseURL = postgresql://postgres:Yipyapyop1@127.0.0.5:5432/$(databaseName)?sslmode=disable
migrationsPath = migrations

build-ui:
	cd ./client && npm install && npm run build

build-server:
	go build main.go;

dev-server:
	go run main.go;

dev-client:
	cd ./client && npm run dev

setup: build-ui build-server

create-migration: 
	migrate create -dir $(migrationsPath) -ext sql $(name)

migrate:
	migrate -database $(databaseURL) -path $(migrationsPath) up 

migrate-down:
	migrate -database $(databaseURL) -path $(migrationsPath) down




