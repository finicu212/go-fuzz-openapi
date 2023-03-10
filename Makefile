build:
	go build -o fuzzctl.exe

test:
	go test

fuzz:
	go test -fuzz Fuzz

server:
	cd third_party/swagger-petstore-swagger-petstore-v3-1.0.17 && mvn package jetty:run

proxy:
	cd tools/proxy-api && npm start