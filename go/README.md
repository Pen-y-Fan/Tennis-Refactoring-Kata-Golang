# GO Starter

This is the refactored version of the Go kata

- Run tests :

```shell
go test ./...
```

- Run tests and coverage :

```shell
go test ./... -coverprofile="coverage.out"

go tool cover -html="coverage.out"
```