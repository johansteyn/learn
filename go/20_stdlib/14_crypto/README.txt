The Golang "crypto" standard library is not FIPS compliant.

To make an application FIPS compliant, using BoringCrypto:
  https://medium.com/cyberark-engineering/navigating-fips-compliance-for-go-applications-libraries-integration-and-security-42ac87eec40b

To build a binary that is not FIPS compliant, use:
  $ go build main.go

To build a binary that is FIPS compliant, use:
  $ GOEXPERIMENT=boringcrypto go build main.go

Note that BoringCrypto is only available on Linux.

