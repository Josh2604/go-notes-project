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
POST /api/v1/notes/create
```
  null
```

**Update error:**
POST /api/v1/notes/update/:id

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