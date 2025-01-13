1. install dependencies
```
go mod tidy
```

2. Compile the protobuf file
   ```
   protoc --go_out=. --go_opt=paths=source_relative spec/apiMessages.proto
   ```
3. This service is a nother service of microservice based application built for demonstration purposes.
4. This service take a number list and sort them in requested method andd send the result with time taken to sort to the api-gateway.
5. See the [medium article](https://medium.com/@kavinduj.20/manage-miroservices-centrally-using-docker-compose-in-both-windows-and-linux-78e61753d284) here for more info.
6. See the service architecture in [here](https://github.com/KMjanith/SR-service-runner/blob/main/Readme.md).
