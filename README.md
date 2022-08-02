# API Spec


## Objects

### Pong message

A pong message is an object which contans 2 fields

* code: int 0 for success, others are failed
* message: string and always be 'pong'

```json
{
  "code": 0,
  "message": "pong"
}
```

### Player

A player object contains 3 fields 

* id: string
* name: string
* age: int

For sample

```json
{
    "id": "1",
    "name": "Ronaldo Cristiano",
    "age": 37
}
```

### Player list

Player list is a collection of players.

```json
[
  {
    "id": "1",
    "name": "Ronaldo Cristiano",
    "age": 37
  },
  {
    "id": "2",
    "name": "De Gea David",
    "age": 31
  },
  {
    "id": "3",
    "name": "Eriksen Christian",
    "age": 30
  },
  {
    "id": "4",
    "name": "Rashford Marcus",
    "age": 24
  },
  {
    "id": "5",
    "name": "Maguire Harry",
    "age": 29
  }
]
```

### Album

```json
 {
    "id": "1",
    "title": "Blue Train",
    "artist": "John Coltrane",
    "price": 56.99
  }
```

### Albums list

```json
[
  {
    "id": "1",
    "title": "Blue Train",
    "artist": "John Coltrane",
    "price": 56.99
  },
  {
    "id": "2",
    "title": "Jeru",
    "artist": "Gerry Mulligan",
    "price": 17.99
  },
  {
    "id": "3",
    "title": "Sarah Vaughan and Clifford Brown",
    "artist": "Sarah Vaughan",
    "price": 39.99
  }
]
```

-------------------------

## Authentications


### Via query string

```
endpoint/?token=secert
```

The endpoint will return 401 if the token is missing or the token is not equal to secert.

### via header

```
token:secert
```
The endpoint will return 401 if the header field is missing or the token is not equal to secert.


----------------

## Endpoints

### GET /ping 

* authenticaton: none
* path: /ping 
* verb：GET
* return value: Pong message

### GET /api/v1/players

* authenticaton: query string
* path: /api/v1/players
* verb：GET
* return value: Player list

### GET /api/v2/players

* authenticaton: header
* path: /api/v2/players
* verb：GET
* return value: Player list

### GET /api/v1/albums

* authenticaton: none
* path: /api/v1/albums
* verb：GET
* return value: Album list

### GET /api/v1/albums/:id

* authenticaton: none
* path: /api/v1/albums/:id
* verb：GET
* return value: An Album Object


### POST /api/v1/albums

* authenticaton: none
* path: /api/v1/albums
* verb：POST
* post data: An Album Object and the id field is required
* return value: The new created Album Object

### DELETE /api/v1/albums/:id

* authenticaton: none
* path: /api/v1/albums/:id
* verb：DELETE
* return value: The new deleted Album Object