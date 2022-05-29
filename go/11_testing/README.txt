To run all the tests from the root folder:

  % go test ./...

Otherwise, change to the subfolder where you want to run tests:

  % cd datastructures

Then, to run all tests in the subfolder, use either of:

  % go test
  % go test .   <= Output to stdout is missing?

For verbose output:

  % go test -v

To run individual tests that match a pattern:

  % go test -run "New"
  % go test -run "10"

To check coverage:

  % go test -coverprofile=c.out
  % go tool cover -html=c.out

The first command above generates the raw data, and the second command processes the data
to generate an HTML page that it displays in the browser.

To run the benchmark tests, use either of:

  % go test -bench=.
  % go test -bench=ListAdd

The above will run all the tests, including benchmark tests.
To run ONLY the benchmark tests, ie. exclude normal tests:

  % go test -run=NONE -bench=.

To run only the benchmark tests that use Add2:

  % go test -run=NONE -bench=ListAdd2

To run the benchmark tests for both Add1 and Add2 100 and 1000 times each:

  % go test -run=NONE -bench=ListAdd._100

To run the benchmark tests for both Add1 and Add2 100 times each:

  % go test -run=NONE -bench=ListAdd._100$

To include memory usage statistics:
  
  % go test -run=NONE -bench=. -benchmem

To produce CPU, memory and block statistics:
  
  % go test -run=NONE -bench=. -cpuprofile=cpu.out
  % go test -run=NONE -bench=. -memprofile=mem.out
  % go test -run=NONE -bench=. -blockprofile=block.out

Running the CPU profiling command above produces the specified profile file as well as file: datastructures.test
Pass those filenames to the pprof tool to produce output:

  % go tool pprof -text -nodecount=10 datastructures.test cpu.out
  % go tool pprof -text -nodecount=10 datastructures.test mem.out
  % go tool pprof -text -nodecount=10 datastructures.test block.out


