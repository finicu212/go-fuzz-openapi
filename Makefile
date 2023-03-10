build:
	go build -o fuzzctl

test:
	go test

fuzz:
	go test -fuzz Fuzz

host:
	cd third_party/swagger-petstore-swagger-petstore-v3-1.0.17 && mvn package jetty:run
