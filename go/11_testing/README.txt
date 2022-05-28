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

Include all the benchmark tests using any of:

  % go test -bench=.
  % go test -bench=ListAdd

To run only the benchmark tests that use add1:

  % go test -bench=ListAdd1

To run the benchmark tests for both add1 and add2 100 and 1000 times each:

  % go test -bench=ListAdd._100

To run the benchmark tests for both add1 and add2 100 times each:

  % go test -bench=ListAdd._100$




