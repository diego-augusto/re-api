# API Challenge
![example workflow](https://github.com/diego-augusto/re-api/actions/workflows/cicd.yml/badge.svg)
[![codecov](https://codecov.io/gh/diego-augusto/re-api/graph/badge.svg?token=XO8CI94YMG)](https://codecov.io/gh/diego-augusto/re-api)



## Running the application

- Make sure `does not have` any other service running on port  or change the `PORT` environment variable.

- See the Makefile for more commands.

#### with go run
```bash
export PORT=8080
make run
```

#### with go docker
```bash
export PORT=8080
make dc-run
```
#### how to call the endpoint
```bash
curl --location 'http://localhost:8080/pack' \
--header 'Content-Type: application/json' \
--data '{
    "items": 12001,
    "sizes" : [5000, 2000, 1000, 500, 250]
}'
```
#### respose
```
[{"size":5000,"quantity":2},{"size":2000,"quantity":1},{"size":250,"quantity":1}]
```


## Routes

`POST /pack`

Description: This endpoint receives a JSON with the items and sizes to be packed.

##### Request

```bash
curl --location 'http://localhost:8080/pack' \
--header 'Content-Type: application/json' \
--data '{
    "items": 12001,
    "sizes" : [5000, 2000, 1000, 500, 250]
}'
```

##### Response

```json
[
    {
        "size": 5000,
        "quantity": 2
    },
    {
        "size": 2000,
        "quantity": 1
    },
    {
        "size": 250,
        "quantity": 1
    }
]
```