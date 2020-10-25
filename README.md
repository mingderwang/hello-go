# hello-go
```
➜  hello-go git:(main) ✗ echo $GOPATH

➜  hello-go git:(main) ✗ go mod init
go: creating new go.mod: module github.com/mingderwang/hello-go
➜  hello-go git:(main) ✗ ls
LICENSE     go.mod      hello.go
README.md   hello       morestrings
➜  hello-go git:(main) ✗ cat go.mod
module github.com/mingderwang/hello-go

go 1.14
➜  hello-go git:(main) ✗ rm -r hello
➜  hello-go git:(main) ✗ go build -o hello hello.go
go: finding module for package github.com/google/go-cmp/cmp
go: found github.com/google/go-cmp/cmp in github.com/google/go-cmp v0.5.2
➜  hello-go git:(main) ✗
```
