# Go API Server for datarepository

Manipulate network conditions via simplified REST calls

## Overview
This server was generated by the [openapi-generator]
(https://openapi-generator.tech) project.
By using the [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) from a remote server, you can easily generate a server stub.
-

To see how to make this your own, look here:

[README](https://openapi-generator.tech)

- API version: 2
- Build date: 2022-05-20T14:51:40.142048+08:00[Asia/Shanghai]


### Running the server
To run the server, follow these simple steps:

```
go run main.go
```

To run the server in a docker container
```
docker build --network=host -t datarepository .
```

Once image is built use
```
docker run --rm -it datarepository
```
