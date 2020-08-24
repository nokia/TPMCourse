# GO-TPM 

Go-TPM is a library written in Golang (or Go) for accessing TPM 2.0 (and 1.2) devices. The current state of the 1.2 libraries is limited so not all functionality is provided and the libraries are not being actively maintained.

The TPM 2.0 libraries however are in good shape and relatively easy to use. It is important to note that that libraries conform to the TPM 2.0 specification and do not try to provide any convenience interfaces - this is a deliberate design decision - which means there are some surprises.

Some familiarity with Go is expected here and not all functionality will be covered. The Go community and the Go-TPM team are pretty friendly however :-)

## Installation

OK, I'm going to assume you have golang installed GOPATH set up and everything works...OK here'a test - save this as `helloworld.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

compile and run with

```bash
go build helloworld.go
./helloworld
```


## Go-TPM
Found here: https://github.com/google/go-tpm

"Go-TPM is a Go library that communicates directly with a TPM device on Linux or Windows machines."

Unlike tpm2tools, GoTPM is a library for communicating with the TPM and supports both TPM 1.2 and 2.0 to some degree - actually TPM 1.2 is depreciated so YMMV. Most of the TPM 2.0 specification is supported.

The library is also very faithful to the specification and provides little in the way of convenience functions. There are some caveats when using the library which need to worked around, but again this is due to the fact that it implements the specification!

So, overall a very useful, functional library for accessing the TPM, but does require some knowledge of the TPM's workings.

## Some Sample Code
In the following sections we show some annotated sample code.

   * gotpm_quote.md
