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
# update file source hello.go
# edit code using package in code
# auto check and -> add missing and remove unused modules
go mod tidy
# path install package [rsc.io/quote]-> /Users/datdao/go/pkg/mod/rsc.io/quote@v1.5.2/quote.go
# Run your code again
go run .

# build -> down load all dependencies and build code
go build

# install package
go get rsc.io/quote
# uninstall package
go get rsc.io/quote@none
# install accounting - money and currency formatting for golang
go get github.com/leekchan/accounting
# update my code
#ac := accounting.Accounting{Symbol: "$", Precision: 2}
#fmt.Println(ac.FormatMoney(123456789.213123)) // "$123,456,789.21"
go build

# func -> https://www.golangprograms.com/go-language/functions.html

# struct
type Calculator struct {
	a int
	b int
}

# import and export package -> https://www.golangprograms.com/golang-import-and-export-struct-packages-and-interfaces/golang-import-struct-from-another-or-child-package.html

# interface -> https://www.youtube.com/watch?v=0Nh8EAlU6eM&list=PLVDJsRQrTUz5icsxSfKdymhghOtLNFn-k&index=6
1. simple interface
2. multiple interfaces
3. embed interface
4. empty interface
```