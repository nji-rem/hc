### Package internal
It contains application logic. Everything is divided into its own domain, so the code is domain partitioned. Keep in 
mind that technically partitioned code will not be permitted. Each module (or component, package) has its own API that
can be found in `api/<module-name>`.

The presentation layer is encouraged to wire dependencies with its corresponding interface, 
e.g. `internal/socket/repository.go` is retrievable by the interface `api/socket/repository.go`.

