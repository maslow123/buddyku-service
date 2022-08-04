## Tools that need to be installed
- [Docker](https://www.docker.com/)
- [Make](https://community.chocolatey.org/packages/make)
- [Go](https://go.dev/)

## Things that must be considered
***Make sure no postgres service is running in the background.**

## How to run the application?

- Open your terminal / cmd, and type the command on below:
    ```
    make runapi
    ```
- Make sure the service is running properly, as follows:
```$ docker ps
CONTAINER ID   IMAGE                                 COMMAND                  CREATED          STATUS          PORTS                      NAMES
0e3c3d387797   maslow123/buddyku-apigateway:latest   "./main"                 31 seconds ago   Up 30 seconds   0.0.0.0:8000->8000/tcp     api-gateway
6495baad4ca4   maslow123/buddyku-articles:latest     "./main"                 34 seconds ago   Up 32 seconds   0.0.0.0:50052->50052/tcp   articleapi
a8722fed4fd2   maslow123/buddyku-users:latest        "./main"                 40 seconds ago   Up 36 seconds   0.0.0.0:50051->50051/tcp   userapi
ccef2920b8d3   postgres:latest                       "docker-entrypoint.s…"   2 minutes ago    Up 2 minutes    0.0.0.0:5433->5432/tcp     testdb
```
- If all services are running well, then import `buddyku.postman_collection.json` into POSTMAN
- Finish

## Documentation API
- You can import the postman collection `buddyku.postman_collection.json` into POSTMAN
- or you can see with swagger command:
    ```make swagger```
- Finish.
## How to run unit testing of each service?
- You just need to enter the command
    ```make test```
- Finish.
## Shut down all services
- Make sure you're on root folder (keuanganku-service), and enter the command
```$ docker-compose down```
- Finish.
