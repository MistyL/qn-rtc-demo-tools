DIR=$(cd ../; pwd)
export GOPATH=$GOPATH:$DIR
go build -o "qlink" main.go