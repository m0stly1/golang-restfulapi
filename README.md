## Golang RESTful API
RESTful API for a public message board

**Docker build**
 - docker build -t golang-restfulapi .
 - docker run -p 8000:8000 -tid golang-restfulapi

**Without Docker**
  - git clone https://github.com/m0stly1/golang-restfulapi.git
  - go get -u ./...
  - go run main.go
  - http://localhost:8000/


# Get Single Message

Get a single message

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

Create a new message with title and content

**URL** : `/message/`

**Method** : `POST`

**Request body**
```json
{
	"title": "A new Hope",
	"content": "A Star Wars title"
}
```
## Success Response

**Code** : `201 OK`

**Content examples**
```
true
```

# Update a message

Update a existing message

**URL** : `/message/{id}`

**Method** : `PUT`

**Request body**
```json
{
	"title": "Force awakens",
	"content": "Another Star Wars title"
}
```

## Success Response
**Code** : `201 OK`

**Content examples**
```
true
```

# Delete a message

Delete a existing message

**URL** : `/message/{id}`

**Method** : `DELETE`


## Success Response
**Code** : `201 OK`

**Content examples**
```
true
```

# Get all messages
Get all messages available in storage

**URL** : `/messages/`

**Method** : `GET`


## Success Response
**Code** : `200 OK`

**Content examples**

```json
{
    "1": {
        "id": "1",
        "title": "Rogue One: A Star Wars Story",
        "content": "More Star Wars"
    },
    "2": {
        "id": "2",
        "title": "Solo: A Star Wars Story",
       	"content": "Almost forgot this one"
    }
}
```

## Responses
| Status Code | Description |
| :--- | :--- |
| 200 | `OK` |
| 201 | `CREATED` |
| 400 | `BAD REQUEST` |
| 404 | `NOT FOUND` |
| 500 | `INTERNAL SERVER ERROR` |