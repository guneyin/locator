# Locator

Simple location marker app to store, list and mark vehicle based location data

## Installation

Clone repository

```bash 
git clone https://github.com/guneyin/locator.git
cd locator
```

Build and run
```bash 
make init
make run
```

Test
```bash 
go test -v ./service/location
```

## API Usage

Go to /api/docs url for Swagger UI
```http
  http://localhost:8081/api/docs
```

## Environment Variables

| Variable       | Desc                       |    Sample |
|----------------|----------------------------|----------:|
| PORT           | HTTP Server Port           |      8081 |
| MAX_RATE_LIMIT | HTTP Server rate limit/sec |        10 |
| DB_HOST        | MySql Database Host        | 127.0.0.1 |
| DB_PORT        | Database Port              |      3306 |
| DB_USER        | Database User              |     mysql |
| DB_PASSWORD    | Database Password          |           |
| DB_NAME        | Database Name              |   locator |

  