# socket-dumper

Dumps everything received on a socket to a file.

## Run the tool

Either compile the tool using the provided `Makefile`. Requires `golang` and `make` to be installed on your system.

```bash
cd src/
make all
```

**Alternatively**, use provided `docker-compose.yaml` file with Docker

```bash
docker-compose up
```

This will start the tool (using port `8000`) and write the recieved socket-data to `out/socket_log.csv`.
