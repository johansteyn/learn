https://levelup.gitconnected.com/using-modules-and-packages-in-go-36a418960556

Packages make use of Go modules, so you need to first run:

  % go mod init packages
  % go mod tidy

Not using modules will result in error:

  % go run packages.go
  packages.go:5:2: package packages/mypackage is not in GOROOT (/usr/local/go/src/packages/mypackage)

TODO: Seems that "go mod tidy" isn't needed?

