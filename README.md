# NoteShare

NoteShare is an application for easily creating and sharing documents of different types online.

# Table of Contents

1. [API specification](#api-specification)

## API specification

The following section contains a specification of the endpoints in the noteshare API.

---

### Get a JSON object representing the structure and content of an account frontpage.

```http
GET /service/api/v1/account/layout
```

#### Description
The structure of the returned JSON object is completely arbitrary, but the client expects a certain structure to be rendered correctly.

#### Responses

##### 200 OK

```http
Content-Type: application/json
```

```json
{
  "layout": null
}
```

##### 401 Unauthorized

Returned if the access token has expired, is invalid or does not exist.

```http
Content-Type: application/json
```

```json
{}
```

##### 404 Not Found

Returned if the account of the associated access token does not exist any more.

```http
Content-Type: application/json
```

```json
{}
```

##### 500 Internal Server Error

```http
Content-Type: application/json
```

```json
{}
```

---

### Update the JSON object representing the structure and content of an account frontpage.

```http
POST /service/api/v1/account/layout
```

#### Responses

##### 200 OK

```http
Content-Type: application/json
```

```json
{}
```

##### 400 Bad Request

Returned if the request body is not valid JSON.

```http
Content-Type: application/json
```

```json
{}
```

##### 401 Unauthorized

Returned if the access token has expired, is invalid or does not exist.

```http
Content-Type: application/json
```

```json
{}
```

##### 415 Unsupported Media Type

```http
Content-Type: application/json
```

```json
{}
```

##### 500 Internal Server Error

```http
Content-Type: application/json
```

```json
{}
```

---

### Get information about the account the access token is a associated with.

```http
/service/api/v1/account/me
```

#### Responses

##### 200 OK

```http
Content-Type: application/json
```

```json
{
  "id": 100,
  "name": "example1"
}
```

##### 401 Unauthorized

Returned if the access token has expired, is invalid or does not exist.

```http
Content-Type: application/json
```

```json
{}
```

##### 404 Not Found

Returned if the account of the associated access token does not exist any more.

```http
Content-Type: application/json
```

```json
{}
```

##### 500 Internal Server Error

```http
Content-Type: application/json
```

```json
{}
```