package helloworld

//go:generate protoc -I.  --go_out=paths=source_relative:. ./helloworld/*.proto
