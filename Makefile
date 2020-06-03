up:
	docker-compose up -d

down:
	docker-compose down

proto:
	docker-compose exec grpc bash -c 'protoc -I ./proto --go_out=plugins=grpc:./proto ./proto/*.proto'
	docker-compose exec grpc bash -c 'protoc -I ./other/proto --go_out=plugins=grpc:./other/proto ./other/proto/*.proto'

mockgen:
	docker-compose exec server bash -c 'mockgen -source domain/model/restaurant/repository.go -destination domain/model/restaurant/mock/mock_repository.go'