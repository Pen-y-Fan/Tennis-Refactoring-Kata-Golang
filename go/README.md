# GO Starter

- Run tests :

```shell
go test ./...
```

- Run tests and coverage :

```shell
go test ./... -coverprofile="coverage.out"

go tool cover -html="coverage.out"
```