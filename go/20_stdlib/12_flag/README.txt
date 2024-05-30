https://pkg.go.dev/flag
https://gobyexample.com/command-line-flags


Build:

  % go build -o flags main.go 


Usage:

  % flags -h
  % flags --h
  % flags -help
  % flags --help


Examples:

1. Single hyphens with equals

  % go run main.go -name=Bob -age=42 -vaccinated=false

2. Double hyphens with equals

  % go run main.go --name=Bob --age=42 --vaccinated=false

3. Single hyphens with spaces

   % go run main.go -name Bob -age 42

   NOTE: Boolean flags cannot use spaces

4. Double hyphens with spaces

  % go run main.go --name Bob --age 42

