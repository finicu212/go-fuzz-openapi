build:
	go test
	go build -o fuzzctl

host:
	cd third_party/swagger-petstore-swagger-petstore-v3-1.0.17 && mvn package jetty:run
