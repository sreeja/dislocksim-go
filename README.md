# dislocksim-go

This is a simulator to measure performance of distributed locks.



#### Testing

- To test
```
go test
```

- Test coverage
```
go test ./... -coverprofile cover.out
go tool cover -func cover.out
```