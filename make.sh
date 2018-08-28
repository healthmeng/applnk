export GOPATH=`pwd`
go get github.com/Go-SQL-Driver/MySQL
go install  -gcflags "-N"  entry
 
