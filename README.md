# TPM 2.0 docker container for experimenting

## Licenses

    The base image for the container is **alpine:latest**.


Software needed for the actual end functionality is the TCG TPM2 Software Stack
(TSS2) and the IBM TPM simulator.

*  [TPM2-tss](https://github.com/tpm2-software/tpm2-tss) BSD-2-Clause
*  [TPM2-abrmd](https://github.com/tpm2-software/tpm2-abrmd) BSD-2-Clause
*  [TPM2-tools](https://github.com/tpm2-software/tpm2-tools) BSD-3-Clause
*  [IBM simulator](./licenses/LICENSE-ibm-tpm-simulator) License included in the
package, extracted here for convenience.

Linux dependecies for building the TCG TPM2 Software Stack are:

* autoconf
* autoconf-archive
* automake
* libtool
* build-base
* pkgconf
* doxygen
* json-c-dev
* openssl
* openssl-dev
* libssl1.1
* git
* udev 
* dbus
* curl-dev
* linux-headers
* glib-dev
* libconfig-dev
* libgcrypt-dev
* wget

It is possible to further strip down dependecies, some packages are only used for
building the documentation which is not in use inside the container as
default.

## Download the container

TODO...will be pushed to dockerhub for download.

## Building the container

Building from source, download repo and run from the repo folder:

    docker build -t tpm2-sim:latest .

Then run with:

    docker run -it tpm2-sim:latest

Remember that the name is different if the image is pulled from the registry.

## TPM 2.0 commands for testing

Let the tpm generate some random data:

    tpm2_getrandom 32 --hex
    
Show the tpm pcr registers:

    tpm2_pcrread
    
Create an endorsment key and make it persistent + check that it is loaded:

    tpm2_createek -c 0x81010001 -G rsa -u ek.pub
    tpm2_readpublic -c 0x81010001

Best documentation for the TPM 2.0 tools can be found at
[TPM software repository](https://github.com/tpm2-software/tpm2-tools/tree/master/man).
