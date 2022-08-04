pull_api: 
	docker-compose down
	docker pull maslow123/buddyku-users
	docker pull maslow123/buddyku-articles
	docker pull maslow123/buddyku-apigateway

infratest: pull_api
	docker-compose up -d --force-recreate testdb
	echo Starting for db...
	sleep 15
	docker-compose up migratedb

runapi: infratest
	docker-compose up -d --force-recreate userapi
	docker-compose up -d --force-recreate articleapi
	docker-compose up -d --force-recreate api-gateway

test: runapi
	cd users && go test -v ./... -coverprofile cover.out
	cd articles && go test -v ./... -coverprofile cover.out
	cd api-gateway && go test -v ./... -coverprofile cover.out
	
	docker-compose down

resetdb:
	docker-compose down
	make infratest

swagger:
	cd api-gateway && make serve-swagger
