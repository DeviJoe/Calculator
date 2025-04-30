# Calculator Swagger API Documentation

This service provides Swagger API documentation for the Calculator REST API.

## Overview

The Swagger UI allows developers to interact with the Calculator API endpoints and understand the expected request/response formats.

## Running the Swagger Server

```bash
go run cmd/calculator_swagger/swagger_server.go
```

By default, the Swagger UI will be available at: http://localhost:8081/swagger

## Configuration

- The default port is 8081
- You can change the port by setting the CALC_SWAGGER_PORT environment variable

## API Endpoints

The API documentation covers:

- `/calc` - POST endpoint for performing calculations

## Available Operations

- Addition (+)
- Subtraction (-)
- Multiplication (*)
- Variable printing