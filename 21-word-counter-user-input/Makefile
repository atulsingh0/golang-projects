
build: counter.go
	env GOOS=darwin GOARCH=amd64 go build -o bin/wc_darwin
	env GOOS=linux GOARCH=amd64 go build -o bin/wc_linux
	chmod u+x bin/wc_*

clean:
	rm -rf bin/wc*
	rmdir bin
