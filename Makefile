export GOPATH=$(PWD)
all:
#	go get github.com/Go-SQL-Driver/MySQL
	rm -f bin/entry
	go install  -gcflags "-N"  entry
 
test:
	go install -gcflags "-N" entry
	cd bin && ./entry

