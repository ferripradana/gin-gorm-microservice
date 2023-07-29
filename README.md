### gin-gorm-microservice

This Code is referenced from https://github.com/gbrayhan/microservices-go

Login API
```
curl --location --request POST 'http://localhost:8080/v1/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"yourmail@gmail.com",
    "password":"yourpassword"
}'
```

Refresh Token
```
curl --location --request POST 'http://localhost:8080/v1/auth/access-token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "refreshToken":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidHlwZSI6ImFjY2VzcyIsImV4cCI6MTY4ODgzMDYzM30.UL-y7LzpuLq3mQMORkYEACZnTXG5qTCz_hP8UuQlK2M"
}'
````

Medicine 
GET ALL with Pagination 
```
curl --location --request GET 'http://localhost:8080/v1/medicine?limit=10&page=1' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidHlwZSI6ImFjY2VzcyIsImV4cCI6MTY4ODg1ODU3N30.j_YDKmaQGxBae1OilFSyr65MlFfgP3Zq-rjEBA4W8nk' \
--header 'Content-Type: application/json' \
```

SWAGGER IS COMING SOON
