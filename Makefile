all: lib/libadd.so lib/libsubtract.so bin/client

lib/libadd.so: add/add.go
	go build -o lib/libadd.so -buildmode=c-shared $<

lib/libsubtract.so: subtract/subtract.go
	go build -o lib/libsubtract.so -buildmode=c-shared $<

bin/client: client.go
	go build -o bin/client client.go
