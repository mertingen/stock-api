# Stock API
It has a RESTful API with two endpoints. One of them that fetches the data in the provided MongoDB collection and returns the results in the requested format. Second endpoint is to create(POST) and fetch(GET) data from an Redis database.

### Assets
\- Go without a web framework
\- MongoDB  
\- Redis

### Warm-up
You'll see **.env.sample** in the root directory, you should fulfill these variable to initialize the app.

DB_USER=
DB_PASS=
DB_HOST=
DB_NAME=

APP_PORT=

REDIS_USER=
REDIS_PASS=
REDIS_HOST=
REDIS_PORT=

### Tests
If'd rather run tests file, you'll enter in "tests" directory and run **go test**

### TO DO
Also, you'll see a **todo_list.md** file in the repository and it lists the matters that the app should have those feature in the future. Because those matters will make the app better.

# 📁 Collection: Records 

## End-point: /api/v1/records
### Method: POST
>```
>localhost:8080/api/v1/records
>```
### Body (**raw**)

```json
{
    "startDate": "2016-01-26",
    "endDate": "2018-02-02",
    "minCount": 2700,
    "maxCount": 3000
}
```

# 📁 Collection: Stocks 


## End-point: /api/v1/records
### Method: POST
>```
>localhost:8080/api/v1/records
>```
### Body (**raw**)

```json
{
    "key": "test-key",
    "value": "zxc-val"
}
```

## End-point: /api/v1/records
### Method: GET
>```
>localhost:8080/api/v1/records?key=test-key
>```
### Query String

```
?key=test-key
```
