# Calculator Service

This project provides a calculator service with both gRPC and REST APIs, along with Swagger documentation.

Algorithm description: [Russian](HOW_IT_WORKS.md) | [English](HOW_IT_WORKS_ENGLISH.md)

SAMPLES [here!](SAMPLE_REQUESTS.md)
## Services

1. **gRPC Server** - Processes calculation requests via gRPC protocol
2. **REST Server** - Provides HTTP access to the calculator service
3. **Swagger UI Server** - Hosts interactive API documentation

## Running with Docker

### Building and Running Individual Services

#### gRPC Server
```bash
docker build -t calculator-grpc -f grpc.Dockerfile .
docker run -p 8081:8081 calculator-grpc
```

#### REST Server
```bash
docker build -t calculator-rest -f rest.Dockerfile .
docker run -p 8080:8080 calculator-rest
```

#### Swagger UI Server
```bash
docker build -t calculator-swagger -f swagger.Dockerfile .
docker run -p 8082:8082 calculator-swagger
```

### Running All Services with Docker Compose

```bash
docker-compose up -d
```

This will start all three services:
- gRPC Server on port 8081
- REST Server on port 8080
- Swagger UI Server on port 8082

## API Usage

### REST API

Access the REST API at `http://localhost:8080/calc`

Example payload:
```json
[
  {
    "type": "calc",
    "var": "x",
    "op": "+",
    "left": "10",
    "right": "5"
  },
  {
    "type": "print",
    "var": "x"
  }
]
```

### Swagger Documentation

Access the Swagger UI at `http://localhost:8082/swagger`

### Environment Variables

- `CALCG_PORT` - Port for the gRPC server (default: 8081)
- `CALCR_PORT` - Port for the REST server (default: 8080)
- `SWG_PORT` - Port for the Swagger UI server (default: 8082)