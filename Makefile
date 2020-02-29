build:
	protoc -I. --go_out=plugins=micro:. \
		proto/vessel/vessel.proto

	docker build -t ryanyogan/shippy-service-vessel:latest .

run:
	docker run -p 50052:50052 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		ryanyogan/shippy-service-vessel:latest