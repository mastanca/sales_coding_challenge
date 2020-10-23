# Sales project challenge proper.ai

## Endpoints

#### `POST` /api/v1/sales

Register a new ticket sale

Expected body:
```json
{
    "country": "AR",
    "event": "lollapalooza"
}
```

Response:

Status: 201, 400, 401
```json
{
    "country": "AR",
    "event": "lollapalooza"
}
```

#### `GET` /api/v1/stats

Get stats of ticket sales by country

Response:

Status: 200, 401
```json
{
    "AR": 2,
    "CL": 1
}
```

#### `POST` /api/v1/login

Log in with user credentials

Expected body:
```json
{
    "username": "testusername",
    "password": "pass"
}
```

Response:

Status: 200, 400, 401
```json
{
    "token": "token"
}
```

## Running

``` shell script
./run.sh
```

## Test
```shell script
go test ./...
```

## Cleanup

``` shell script
./clean.sh
```

## Credentials

As saving users is out of scope, a simple JWT approach was implemented
with a fake hardcoded user saved.
```
username: testusername
password: pass
```


