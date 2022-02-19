Run:

  % go run helloworld.go

Build executable:

  % go build helloworld.go

Install binary:

  % go install helloworld.go

This will put the "helloworld" binary in the $GOPATH/bin folder:

  ~/go/bin/01_helloworld

NOTE: When using modules you can simply run:

  % go build
  % go install

Now, with this being on Github, if I run:

  % go get github.com/johansteyn/learn/go/01_helloworld

Then it will download all of my "learn" repository to:

  ~/go/pkg/mod/github.com/johansteyn/learn@v0.0.0-20220206092148-10bea03d5acf

Then, if I run:

  % go install github.com/johansteyn/learn/go/01_helloworld@latest

It will build my app and place the binary in the $GOPATH/bin folder:

  ~/go/bin/01_helloworld

Note that the name of the binary includes "01_", which is different from running install earlier.
I assume go takes the folder name as the default name of the module when modules are not being used.

