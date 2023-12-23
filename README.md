# Spam47 Service API

Spam47 is a service that allows you to check if a given message is spam or not (ham), and train the service with new data.

## API Endpoints

### Health Check

`GET /livez`

Returns the status of the service.

### Check Message

`POST /check`

Checks if a given message is spam or not.

Request body:

```json
{
  "message": "Hello world",
  "lang": "en"
}
```

Response:

```json
{
  "status": "spam",
  "score": 0.9995
}
```

### Train Service

`POST /train`

Trains the service with a new message.

Request body:

```json
{
  "message": "Hello world",
  "type": "ham",
  "lang": "en"
}
```
