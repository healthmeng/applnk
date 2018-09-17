export GOPATH=$(PWD)
all:
	go get github.com/Go-SQL-Driver/MySQL
	go install  -gcflags "-N"  entry
 
test:
	go install -gcflags "-N" entry
	cd bin && ./entry

