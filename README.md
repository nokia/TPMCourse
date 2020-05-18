# TPM 2.0 docker container for experimenting

Software needed for the actual end functionality is the TCG TPM2 Software Stack
(TSS2) and the IBM TPM simulator.

*  [TPM2-tss](https://github.com/tpm2-software/tpm2-tss) BSD-2-Clause
*  [TPM2-abrmd](https://github.com/tpm2-software/tpm2-abrmd) BSD-2-Clause
*  [TPM2-tools](https://github.com/tpm2-software/tpm2-tools) BSD-3-Clause
*  [IBM simulator](./licenses/LICENSE-ibm-tpm-simulator)

## Download the container

To be added on DockerHub...

## Building the container

Building from source, download repo and pick distribution. Run from the distro folder:

    docker build -t tpmcourse:latest .

Then run with:

    docker run -it tpmcourse:latest

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


# Getting started with TPM

Under [./docs](./docs), you can find a set of files that explain different TPM concepts and contain code examples you can use with the container to get familiarized with TPM and its use cases.