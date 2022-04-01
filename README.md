# GO notes - course project

### Requirements
* Install go.mod dependencies

### Configs:
**config.yml**: Contains env variables
**init.go**: Initialize variables from config file
```txt
config
├── config.yml
└── init.go
```
### Run locally
**Install make to run commands**
```sh
# for MacOS
brew install make
```

Run app locally using the next commands:
**Build postgres container**
```
make build-pg
```
**Run postgres container**
```
make run-postgres
```
If you want use mongo service:
```
make run-mongo
```
### Run app
```
make run
```

### Run with Docker
**Note: You must comment DB implementations to run**
```
docker build --pull --rm -f "Dockerfile" -t gonotesproject:latest "."
docker run -it -p 8080:8080 gonotesproject:latest
```
### Enpoints

**Health:**
GET /api/v1/notes/ping

```txt
  pong
```

**Create note:**
POST /api/v1/notes

***request:***
```json
  {
      "name": "{note_name}",
      "description": "{note_description}"
  }
```
***response:***
```json
  {
      "code": 201,
      "message": "Resource created successfully"
  }
```

**Update error:**
PUT /api/v1/notes/:id

***request:***
```json
  {
      "name": "Name updated"
  }
```
***response:***
```json
  {
      "code": 200,
      "message": "Resource updated successfully"
  }
```

**Get All:**
GET /api/v1/notes/all

***response:***
```json
  {
      "code": 200,
      "message": [
          {
              "ID": "2",
              "Name": "Test name 2",
              "Description": "Nota de test 2",
              "Deleted": false,
              "DateCreated": "2021-10-02T22:56:15.105786Z",
              "DateUpdated": "0001-01-01T00:00:00Z",
              "DateDeleted": "0001-01-01T00:00:00Z"
          },
          {
              "ID": "3",
              "Name": "Note test 4",
              "Description": "This is the default description",
              "Deleted": false,
              "DateCreated": "2021-10-02T18:06:22.263068Z",
              "DateUpdated": "0001-01-01T00:00:00Z",
              "DateDeleted": "0001-01-01T00:00:00Z"
          },
          {
              "ID": "1",
              "Name": "Test name 1",
              "Description": "Updated note",
              "Deleted": false,
              "DateCreated": "2021-10-02T22:56:15.105786Z",
              "DateUpdated": "2021-10-02T18:07:14.568577Z",
              "DateDeleted": "0001-01-01T00:00:00Z"
          }
      ]
  }
```

**Get One Note:**
GET /api/v1/notes/:id

***response:***
```json
{
    "ID": "1",
    "Name": "Test name 1",
    "Description": "Nota de test 1",
    "Deleted": true,
    "DateCreated": "2021-09-25T00:19:30.388258Z",
    "DateUpdated": "0001-01-01T00:00:00Z",
    "DateDeleted": "2021-09-25T00:26:03.811343Z"
}
```

**Delete Note:**
DELETE /api/v1/notes/:id

***response:***
```json
  {
      "code": 200,
      "message": "Resource deleted successfully"
  }
```

**SignUp:**
POST /api/v1/notes/signup

***request:***
```json
  {
      "username": "{user_name}",
      "password": "{password}"
  }
```
***response:***
```json
  {
      "code": 201,
      "message": "User registered successfully"
  }
```

**SignIn:**
POST /api/v1/notes/signin

***request:***
```json
  {
      "username": "{user_name}",
      "password": "{password}"
  }
```
***response:***
```json
  {
      "token": "{token}"
  }
```