TARGETS = \
	socket-dumper \
	socket-dumper-win64.exe \
	socket-dumper-win32.exe \
	tick2socket \
	tick2socket-win64.exe \
	tick2socket-win32.exe \


all: ${TARGETS}

clean:
	touch ${TARGETS}
	rm    ${TARGETS}

mrproper: clean
	touch dummy~
	rm        *~


socket-dumper: socket-dumper.go
	go build socket-dumper.go

socket-dumper-win64.exe: socket-dumper.go
	GOOS=windows GOARCH=amd64 go build -o socket-dumper-win64.exe socket-dumper.go

socket-dumper-win32.exe: socket-dumper.go
	GOOS=windows GOARCH=386 go build -o socket-dumper-win32.exe socket-dumper.go

tick2socket: tick2socket.go
	go build tick2socket.go

tick2socket-win64.exe: tick2socket.go
	GOOS=windows GOARCH=amd64 go build -o tick2socket-win64.exe tick2socket.go

tick2socket-win32.exe: tick2socket.go
	GOOS=windows GOARCH=386 go build -o tick2socket-win32.exe tick2socket.go

