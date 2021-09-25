# GO notes - course project

### Requirements
* Install go.mod dependencies

### Configs:
**config.yml**: Contains env variables
**init.go**: Initialize variables from config file
```
config
├── config.yml
└── init.go
```


### Enpoints

**Health:**
GET /api/v1/notes/ping

```
  pong
```

**Create note:**
POST /api/v1/notes
```
  null
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
```
  null
```

**Get All:**
GET /api/v1/notes/all

***response:***
```
  null
```

**Get One Note:**
GET /api/v1/notes/:id

***response:***
```
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
```
  null
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
```
  null
```