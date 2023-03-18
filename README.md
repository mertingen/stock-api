# Stock API
It has a RESTful API with two endpoints. One of them that fetches the data in the provided MongoDB collection and returns the results in the requested format. Second endpoint is to create(POST) and fetch(GET) data from an Redis database.

### Assets
- Go without a web framework
- MongoDB
- Redis

### Warm-up
You'll see **.env.sample** in the root directory, you should fulfill these variable to initialize the app.

- DB_USER=
- DB_PASS=
- DB_HOST=
- DB_NAME=

- PORT=

- REDIS_USER=
- REDIS_PASS=
- REDIS_HOST=
- REDIS_PORT=

### Tests
If'd rather run tests file, firstly you need to fufill **.env.test** file while checking **.env.sample** file and you'll enter in "tests" directory then run **go test**

### TO DO
Also, you'll see a **todo_list.md** file in the repository and it lists the matters that the app should have those feature in the future. Because those matters will make the app better.

# 📁 Collection: Records 

The request payload of the first endpoint will include a JSON with 4 fields.
● “startDate” and “endDate” fields will contain the date in a “YYYY-MM-DD” format.
You should filter the data using “createdAt”
● “minCount” and “maxCount” are for filtering the data. Sum of the “count” array in
the documents should be between “minCount” and “maxCount”.

Response payload should have 3 main fields.
● “code” is for status of the request. 0 means success. Other values may be used
for errors that you define.
● “msg” is for description of the code. You can set it to “success” for successful
requests. For unsuccessful requests, you should use explanatory messages.
● “records” will include all the filtered items according to the request. This array
should include items of “key”, “createdAt” and “totalCount” which is the sum the
“counts” array in the document.

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

The request payload of POST endpoint will include a JSON with 2 fields.

● “key” fields holds the key (any key in string type)
● “value” fields holds the value (any value in string type)

Response payload should return echo of the request or error (if any).

The request payload of GET endpoint will include 1 query parameter. That is “key”
param holds the key (any key in string type)

## End-point: /api/v1/stocks
### Method: POST
>```
>localhost:8080/api/v1/stocks
>```
### Body (**raw**)

```json
{
    "key": "test-key",
    "value": "test-val"
}
```

## End-point: /api/v1/stocks
### Method: GET
>```
>localhost:8080/api/v1/stocks?key=test-key
>```
### Query String

```
?key=test-key
```
