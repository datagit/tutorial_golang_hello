https://golang.org/doc/tutorial/getting-started
```bash
# check go version
go version
mkdir hello
cd hello

# Enable dependency tracking for your code.
go mod init example.com/hello
# Run your code to see the greeting.
go run .
# help
go help
# add missing and remove unused modules
go mod tidy
# path install package [rsc.io/quote]-> /Users/datdao/go/pkg/mod/rsc.io/quote@v1.5.2/quote.go
# Run your code again
go run .
```