build:
	echo "Building proto buffer file"

	protoc -I. --go_out=plugins=micro:. \
		proto/vessel/vessel.proto

	echo "Building docker image"

	docker build -t ryanyogan/shippy-service-vessel:latest .

run:
	echo "Running docker image"

	docker run -p 50052:50052 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		ryanyogan/shippy-service-vessel:latest