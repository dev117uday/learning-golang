# Learning Golang

Directory Structure
```
.
├── boilerplate-go
├── built-in-dsa-go
├── competitve-go
├── concurrency_patterns
├── ds-go
├── go-book
├── go.mod
├── ide-go
├── miscellaneous
├── programs-go
├── README.md
└── tour-go
```


## Golang tooling

Video : [Go tooling](https://www.youtube.com/watch?v=uBjoTxosSys)

- Formattor : `gofmt main.go`
    -  to view diff use `-d`

- to get all import inside a program : `goimports`
- to compile for windows : `GOOS=windows go build`
- to install a program : `go install` -> it installs it in GOPATH

### Go List :

- to list the path of package : `go list`\
- to print the name of the file : `go list -f '{{ .Name }}'`
- to print the documentation of the program : `go list -f '{{ .Doc }}' `
- to print all the imports : `go list -f '{{ .Imports }}'`
- to get all packages that one of the package depends on : `go list -f '{{ join .Imports "\n"  }}' fmt`
- to get documentation for package : `go doc fmt` (using "fmt" as example)
- to get doc for a function inside a package : `go doc fmt Println` (using "fmt" as package and "Println" as function)
- to start documentation server : `godoc -http :6060`
- to check for possible errors in program : `go vet`


