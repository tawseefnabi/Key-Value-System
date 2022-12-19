# Highly-Available-Key-Value-System
## Description:
Built scalable key-value system, which allow user to add, put and delete value at high scale

## Architecture(HLD):
<img width="712" alt="Screenshot 2022-12-10 at 6 11 53 PM" src="https://user-images.githubusercontent.com/39022530/206855918-f505e0e7-ce9f-4e5d-ad71-6890df4a3c62.png">

This system is designed in such a way that, it can scale up to million of user. I have divide my project into two part, that is cron job and main application, It will allow user to create, update, get and delete the data at high low latency

1. Cron Job:
This compoment is responsible for cleaning up the data from the database, which are already expired or marked as soft-delete.

2. KeyValueMain:
This compoment is responsible for performing operation of create, update, delete and get on database. 

## Tech Stack:
1. golang
2. cron job

## Set Up:
1. Install golang

## How to run project?

1. Invoke KeyValueCron:
```
  go run .
```

2. Invoke KeyValueMain:
```
  go run .
```

## Requests:
1. Put Data
```
curl --location --request PUT '127.0.0.1:8000/api/put' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key":"key1",
    "value":"killer4",
    "ttl":0
}'
```
2. Get Data:
```
curl --location --request GET '127.0.0.1:8000/api/get?key=key1'
```
3. Delete Data:a
```
curl --location --request DELETE '127.0.0.1:8000/api/delete?key=key1'
```
