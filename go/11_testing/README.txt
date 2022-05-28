To run all tests in this folder:

  % go test

For more verbose output:

  % go test -v

To run individual tests that match a pattern:

  % go test -run "With"

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


To run only the benchmark tests that use add1:

  % go test -run=NONE -bench=ListAdd1

To run the benchmark tests for both add1 and add2 100 and 1000 times each:

  % go test -run=NONE -bench=ListAdd._100

To run the benchmark tests for both add1 and add2 100 times each:

  % go test -run=NONE -bench=ListAdd._100$

To include memory usage statistics:
  
  % go test -run=NONE -bench=. -benchmem

To produce CPU, memory and block statistics:
  
  % go test -run=NONE -bench=. -cpuprofile=cpu.out
  % go test -run=NONE -bench=. -memprofile=mem.out
  % go test -run=NONE -bench=. -blockprofile=block.out

Running the CPU profiling command above produces file: test_list.test and cpu.out
Pass those filenames to the pprof tool to produce output:

  % go tool pprof -text -nodecount=10 test_list.test cpu.out
