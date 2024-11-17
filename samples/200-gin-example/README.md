# Gin example

```shell
go mod init 200-gin-example
go get github.com/gin-gonic/gin@latest
```

## Test

```shell
curl --request GET \
  --url 'http://localhost:8080/hello/message?test=someTestValue' 
# Response:
#{
#	"test": "better someTestValue"
#}
```

```shell
curl --request POST \
  --url http://localhost:8080/hello/message \
  --header 'Content-Type: application/json' \
  --data '{
	"message1": "Test message 1",
	"message2": "Test message 2"
}'
# Response:
#{
#	"answerMessage": "Test message 1"
#}
```

## Docs

https://pkg.go.dev/github.com/gin-gonic/gin
