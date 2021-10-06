package main

import (
    "flag"
	"fmt"
    "reflect"

	"github.com/google/go-tpm/tpm2"
	"github.com/google/go-tpm/tpmutil"
)


var (
	tpmPath = flag.String("tpm-path", "/dev/tpm0", "Path to the TPM device (character device or a Unix socket)")
)

func main() {

	rwc, err := tpm2.OpenTPM(*tpmPath)
	if err != nil {
		fmt.Errorf("can't open TPM at  %v", err)
	}
	//defer rwc.Close()

        nonce := []byte{1, 2, 3, 4, 5, 6, 7, 8}

        var handle = tpmutil.Handle(0x810100AA)

        fmt.Println("Handle is",handle)
    	fmt.Println("Type:",reflect.TypeOf(handle).String())

        // quote retuns 3 values,  att, sig, err
        att, sig, err := tpm2.Quote(
                 rwc,
                 handle,
 		"",
		"",
		nonce,
                tpm2.PCRSelection{ tpm2.AlgSHA256 , []int{0} },
		tpm2.AlgNull)

	if err != nil {
		fmt.Errorf("Problem getting quote  %s", err)
                fmt.Println(err)
	}
	if sig != nil {
		fmt.Errorf("Sig is nil\n")
        }
	if att != nil {
		fmt.Errorf("Att is nil\n")
        }
        // att is of type []byte
        // sig is of type tpm2.Signature ??
        // err is ??
        fmt.Println("Err is *",err,"*")
        fmt.Println("length  ",len(att))
        fmt.Println("Att [% x] ",att)
    

}
