# codeid

- create murid
```
curl --location 'localhost:8082/api/murid' \
--header 'Content-Type: application/json' \
--data '{
    "name":"test",
    "id": 1
}'
`
```

- get murid 
```
curl --location 'localhost:8082/api/murid'
```
-- update murid
```
curl --location --request PUT 'localhost:8082/api/murid/1' \
--header 'Content-Type: application/json' \
--data '{
    "name": "wow",
    "id":1
}'
```