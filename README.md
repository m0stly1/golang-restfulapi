## Golang RESTful API
RESTful API for a public message board

**Docker build**
 - docker build -t name
 - docker run -p 8000:8000 -tid name

**Without Docker**
  - git clone https://github.com/m0stly1/playground1.git
  - go get -u ./...
  - go run main.go
  - http://localhost:8080/


# Get Single Message

Get the details of the currently Authenticated User along with basic
subscription information.

**URL** : `/message/{id}`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Content examples**

```json
{
	"id": "1",
	"title": "A new Hope",
	"Content": "Title got a star wars reference"
}
```

# Create a new message

Get the details of the currently Authenticated User along with basic
subscription information.

**URL** : `/message/`

**Method** : `POST`

**Request body**
```json
{
	"title": "A new Hope",
	"content": "A Star Wars title"
}
```


**Code** : `201 OK`

**Content examples**
```
true
```

# Update a message

Get the details of the currently Authenticated User along with basic
subscription information.

**URL** : `/message/{id}`

**Method** : `PUT`

**Request body**
```json
{
	"title": "Force awakens",
	"content": "Another Star Wars title"
}
```


**Code** : `201 OK`

**Content examples**
```
true
```



## Responses
| Status Code | Description |
| :--- | :--- |
| 200 | `OK` |
| 201 | `CREATED` |
| 400 | `BAD REQUEST` |
| 404 | `NOT FOUND` |
| 500 | `INTERNAL SERVER ERROR` |