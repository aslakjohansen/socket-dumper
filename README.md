# socket-dumper

Dumps everything received on a socket to a file (along with timestamps).

## Build Dependencies

- `golang`
- `make`

## Building

To compile:

```shell
cd src/ ; make
```

## Running

Run on port 8000 and log to `socket_log.csv`:

```shell
cd src/ ; ./socket-dumper 8000 socket_log.csv
```

## Docker Wrapper

Use provided `docker-compose.yaml` file with Docker:

```shell
docker-compose up
```

This will start the tool (using port `8000`) and write the recieved socket-data to `out/socket_log.csv`.

