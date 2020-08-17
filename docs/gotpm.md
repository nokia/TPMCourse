# GO-TPM 

Go-TPM is a library written in Golang (or Go) for accessing TPM 2.0 (and 1.2) devices. The current state of the 1.2 libraries is limited so not all functionality is provided and the libraries are not being actively maintained.

The TPM 2.0 libraries however are in good shape and relatively easy to use. It is important to note that that libraries conform to the TPM 2.0 specification and do not try to provide any convenience interfaces - this is a deliberate design decision - which means there are some surprises.

Some familiarity with Go is expected here and not all functionality will be covered. The Go community and the Go-TPM team are pretty friendly however :-)

## Installation and First Example