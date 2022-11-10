build-ui:
	cd ./client && npm install && npm run build

build-server:
	go build main.go;

dev-server:
	go run main.go;

dev-client:
	cd ./client && npm run dev

setup: build-ui build-server



