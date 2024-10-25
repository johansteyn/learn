
The "dialer" client will not allow local connections.

For example, these connections will not be allowed:

  % go run dialer.go "localhost" 
  % go run dialer.go "127.0.0.1"
  % go run dialer.go "127.0.0.53"
  % go run dialer.go "0.0.0.0"

Whereas these will be allowed:

  % go run dialer.go "example.com"
  % go run dialer.go "93.184.215.14"  <= The IP address of example.com
  
