https://levelup.gitconnected.com/using-modules-and-packages-in-go-36a418960556
https://tip.golang.org/doc/go1.16#modules
https://www.youtube.com/watch?v=Z1VhG7cf83M
https://www.youtube.com/watch?v=20sLKEpHvvk <= Very good

A Go application can use local packages and depend on 3rd party code.

Whereas a package is a collection of source files that are compiled together,
a module is a collection of packages that are released together.

If we have "mymodule" with packages "mypackage1" and "mypackage2", then
they can be imported as: "mymodule/mypackage1" and "mymodule/mypackage2"

Not using modules will result in error:

  % go run myapp.go        
  myapp.go:5:2: package mymodule/mypackage1 is not in GOROOT (/usr/local/Cellar/go/1.17.6/libexec/src/mymodule/mypackage1)
  myapp.go:6:2: package mymodule/mypackage2 is not in GOROOT (/usr/local/Cellar/go/1.17.6/libexec/src/mymodule/mypackage2)
  myapp.go:7:2: no required module provides package github.com/hackebrot/turtle: go.mod file not found in current directory or any parent directory; see 'go help modules'

Before Go version 1.15 and modules, all local packages had to be placed in 
$GOROOT or $GOPATH/src, and you needed to run "go get" to download 3rd party
dependencies to $GOPATH/src:

  % go get github.com/hackebrot/turtle

It would then run "go install" to compile the downloaded 3rd party sources
and place them in $GOPATH/pkg.

With modules you no longer need to run "go get" or "go install".

To create a module, generate a go.mod file:

  % go mod init mymodule

The go.mod file declares the "module path", which is the import path prefix
for all packages within the module.

The go.mod file also lists the module's dependencies, but they won't have been 
added during "init" as it won't be able to figure out the dependencies until 
the source code has been written...

So, after writing your source code, add module dependencies and sums with:

  % go mod tidy

Note that "tidy" is only required if your module has dependencies.
It will parse your source code to determine the dependencies, add them
to your go.mod file and then download and place them in $GOPATH/pkg/mod.
ie. "go mod tidy" effectively does what "go get" used to do.

Interesting... if you run "go get", eg:
  % go get rsc.io/quote
Then it will add to go.mod and go.sum without looking at your source code.
If you decide not to use those added dependencies in your source code, 
then running "tidy" will remove those added entries.
I cannot see any reason why you would run "go get" on its own when using modules...

NOTE:
Go will use the go.mod file in the current folder, or any parent folder.

Question:
Should you keep go.mod and go.sum files under version control?
Yes: https://github.com/golang/go/wiki/Modules

NOTE:
Learned later that when developing a new app, you can start coding and run "go mod tidy" 
to add dependencies to the go.mod file. But if you already know that you will depend on
a particular library before even writing a line of code, you can add that dependency to
the go.mod file by running "go get", eg. when I was getting ready to write an app that
would use go-redis:
  % go mod init mymodule
  % go get github.com/redis/go-redis/v9
However, if you then run "go mod tidy" it will remove the just-added dependencies,
obviously since it will not find any code that uses it.
So, "go get" is actually still used, if you want to.
But note that the go.mod file produced by "go get" is somewhat different from that 
produced by "go mod tidy"...
So, I will prefer to use "go mod tidy" over "go get".


