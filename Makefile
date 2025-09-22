run:
	docker build -t warehouse-mgt-app . && docker image prune -f && docker run -it --rm \
  	--network app-network \
  	-p 8334:8334 \
  	-v $(pwd)/config.json:/config.json \
  	warehouse-mgt-app


build:
	go build -o app cmd/app/main.go