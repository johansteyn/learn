https://github.com/spf13/cobra

Using Cobra to develop a simple "lfs" program to run filesystem commands, such as:

  - List the contents of directories
  - Create new directories
  - Remove files and directories

Following the Cobra pattern of:

  APPNAME COMMAND ARG --FLAG

Examples:

  % fs list <dir>
  % fs list <dir> -l
  % fs list <dir> --long


First, install Cobra:

  % go get -u github.com/spf13/cobra

This will place the Cobra and PFlags packages in:

  ~/go/pkg/mod/github.com/spf13


Then, install the Cobra binary:

  % go get github.com/spf13/cobra/cobra

This will place the "cobra" binary in:

  ~/go/bin

Use the "cobra" command to create and initialize an application:

  % ~/go/bin/cobra init fs --pkg-name fs
 
Use "go mod" to create a module:

  % cd fs
  % go mod init fs
  % go mod tidy

Run the app:

  % go run main.go

Renamed the main source file so that the binary is called "fs", and built the binary:

  % mv main.go fs.go
  % go build fs.go

To add a command to the application:

  % ~/go/bin/cobra add list

This will create a new file:

  ~/cmd/list.go

The "init" function is where you define flags and config settings.

Added a "--long" flag to the "list" command, so it can be used as:

  % ./fs list -l
  % ./fs list --long
  % ./fs list --long=true
  % ./fs list --long=false

Manually added a "cd" command by copying cmd/list.go to cmd/cd.go (and editing)
The "cd" command has no flags, but it requires at least one argument.

Added an "rm" command, with a required "force" flag (and also at least one argument).

NOTE: Even though the "main" function is the normal entry point, each file's "init" function is called first.
  https://tutorialedge.net/golang/the-go-init-function/
The order of "init" function calls across files is not specified - so don't rely on any order!

