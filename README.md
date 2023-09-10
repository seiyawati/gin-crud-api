# Gin CURD API

## 環境構築

```
docker-compose up -d --build
```

## Hot Reload

AirによってHot Reloadされるので再コンパイルは不要
```
app_1  |
app_1  |   __    _   ___
app_1  |  / /\  | | | |_)
app_1  | /_/--\ |_| |_| \_ , built with Go
app_1  |
app_1  | mkdir /go/src/app/tmp
app_1  | watching .
app_1  | watching controllers
app_1  | !exclude data
app_1  | watching models
app_1  | watching requests
app_1  | !exclude tmp
```

## API

### GET /todos

```
curl -X GET http://localhost:8080/todos
```

### GET /todos/:id

```
curl -X GET http://localhost:8080/todos/:id
```

### POST /todos

```
curl -X POST -H "Content-Type: application/json" -d '{"title":"test", "author":"test"}' http://localhost:8080/todos
```

### PATCH /todos/:id

```
curl -X PATCH -H "Content-Type: application/json" -d  '{"title":"test", "author":"test"}' http://localhost:8080/todos/1
```

### DELETE /todos/:id

```
curl -X DELETE http://localhost:8080/todos/1
```
