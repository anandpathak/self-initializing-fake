# Self-initializing-fake
A fake server for mock testing and contract testing.

This is inspired from one of the Martin Fowler's blog https://martinfowler.com/bliki/TestDouble.html

### Installation

Run via docker

```
 docker pull anandpathak/self-initializing-fake
 docker run -p 8112:8112 -p 8113:8113 self-initializing-fake:latest
```



## [API defination](doc/swagger.yaml)

### Example setup fake route

```
curl --location --request POST 'localhost:8112/setup/fake_route' \
--header 'Content-Type: application/json' \
--data-raw '{
    "request": {
        "body": {
            "a": 1
        },
        "headers": {
            "Test": ["abc","xyz"]
        }
    },
    "response": {
        "headers": {
            "Test": ["test"],
            "Content-Type": ["application/json"]
        },
        "body": {
            "abc": "awesome"
        }  
    },
    "url": "/test1"
}'
```
### Call fake route
```
 curl --location --request POST 'localhost:8113/test' \
--header 'test: abc' \
--header 'Content-Type: application/json' \
--data-raw '{
    "a":1
}'
```

