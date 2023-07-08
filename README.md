### gin-gorm-microservice

This Code is referenced from https://github.com/gbrayhan/microservices-go

Login API
```
curl --location --request POST 'http://localhost:8080/v1/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"ferri.pradana@gmail.com",
    "password":"Password"
}'
```

