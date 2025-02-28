# Telenotify
Golang telegram bot to send you notification that you got from online services via grpc or http calls.
Its like simple webhook for you to get notifications in telegram from web.
Using postgres docker image to store subscribers.

## Run in docker
`docker compose up -d`

Use `/ping` command in chat with bot to check if bot up and running

## Get notinifcations

Use `/subscribe` command in chat with bot

## Send notifications

Use grpc 

```
message NotifyRequest {
  string message = 1;
  string sign = 2;
}
```
or http post call
```
{
  "message": "notification message"
  "sign": "secret"
}
```

## Security

Checking sign secret not added yet