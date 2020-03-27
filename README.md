# fortune
A web server to obtain Unix fortune messages (for fun)!

## Building and running
You can build this project using Docker:
```
docker built -t fortune:latest .
docker run -p 8080:8080 fortune:latest
```

You can also build this without Docker:
```
go build
go run main.go
```

You can try hitting the server once it's running using any of the following:
```
curl 'http://127.0.0.1:8080/'
curl 'http://127.0.0.1:8080/cookie'
curl 'http://127.0.0.1:8080/offensive-cookie' 
```

## Testing
Honestly, not much to test, but you could run tests using:
```
go test
```
