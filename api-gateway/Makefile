proto:
	protoc \
		-I D:/me/exercise/go/pkg/mod/github.com/gogo/protobuf@v1.3.2 \
		-I D:/me/exercise/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto \
		-I pkg pkg/users/pb/*.proto --gofast_out=plugins=grpc:. 
	protoc \
		-I D:/me/exercise/go/pkg/mod/github.com/gogo/protobuf@v1.3.2 \
		-I D:/me/exercise/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto \
		-I pkg pkg/articles/pb/*.proto --gofast_out=plugins=grpc:. 

buildapigateway:
	docker-compose build api-gateway

runapigateway: buildapigateway
	docker-compose up -d api-gateway

server:
	go run cmd/main.go

test:
	go test -v ./... -coverprofile cover.out

check-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	swagger generate spec -o ./docs/swagger.yaml --scan-models

serve-swagger: swagger
	swagger serve -F=swagger ./docs/swagger.yaml