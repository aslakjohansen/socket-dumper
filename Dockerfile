FROM golang:1.22-bullseye

WORKDIR /src

COPY src/ .

RUN make socket-dumper

RUN mkdir /out

ENTRYPOINT ./socket-dumper 8000 /out/socket_log.csv