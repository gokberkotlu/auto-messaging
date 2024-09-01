# AUTO MESSAGING

This application sends unsent message data from the database to an external API at specific intervals and caches the message IDs and timestamps. It reads messages from a dataset and performs batch-load to the database. To avoid repeating the batch-load, it performs a migration check.

## Endpoints

- Start/Stop automatic message sending
- Retrieve a list of sent messages

[Swagger Link](http://localhost:8080/swagger/index.html)

## Environments

It is necessary to create a .env file for the Docker containers used by the application and the environment variables used within the application.

### Example .env File

```bash
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=auto_messaging_db
POSTGRES_PORT=54321
REDIS_PASSWORD=redis
REDIS_PORT=63791
CLIENT_TOKEN=
CLIENT_X_INS_AUTH_KEY=
```

## Containers

To run the containers

```bash
docker-compose up -d
```

## Run Command to Start the Application

```bash
go run main.go
```

# Run Command to Update Swagger Docs

```bash
./swag.sh
```
