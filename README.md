
### Tech Stack ###
* GO
* Fiber
* MYSQL
* HTTP Only cookies
* JWT

## Perintah untuk menjalankan 
```sh
go run main.go
```

## Routes
```sh
## AUTH
localhost:8000/api/register # Register(s) (POST)
localhost:8000/api/login # Login user (POST)
localhost:8000/api/logout  # Logout user (POST)
localhost:8000/api/user # get user (get)

## CRUD
localhost:8000/api/create # Menambahkan notes (POST)
localhost:8000/api/note/:id # get notes by id (GET)
localhost:8000/api/notes # get all notes (GET)
localhost:8000/api/update/:id # update notes by id (PUT)
localhost:8000/api/delete/:id # delete notes by id (DELETE)
```

## Sample Request in postman
##### Register 
```json
{
    "name": "emon",
    "email": "emon@gmail.com",
    "password": "8798634"

}
```

##### Login 
```json
{
    "email": "emon@gmail.com",
    "password": "8798634"
}
```

##### Menambahkan Notes 
```json
{
    "title": "test",
    "category": "test pertama",
    "details": "data"
}

```

##### Update notes by id
```json
{
    "title": "test",
    "category": "test pertama",
    "details": "data"
}

```
