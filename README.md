# Nokia TPM Course

This is a short course on getting started with understanding how a TPM 2.0 works. In this course we explain a number of the features of the TPM 2.0 through the TPM2_Tools through examples and, optionally, exercises.

## Installation

Or actually, how to run the docker containers which provide the tools and the IBM TPM Simulator so you can safely practice without running against your or someone else's real TPM.


We provide a Dockerfile for the Linux Alpine distribution which when built downloads the various libraries and tools to interact with a TPM and also the IBM TPM Simulator

```bash
cd alpine
docker build -t tpmcourse:latest .
```

If all works successfully, after about 29 stages, you can type the following command to run the container in interactive mode.

```bash
docker run -it tpmcourse:latest
```

The following will be printed out. This tells us that the TPM simulator is runnig on ports 2321 and 2322 and at the end we have a command line prompt. We're running the `ash` command line interpreter.

```
/ # LIBRARY_COMPATIBILITY_CHECK is ON
Manufacturing NV state...
Size of OBJECT = 2600
Size of components in TPMT_SENSITIVE = 1096
    TPMI_ALG_PUBLIC                 2
    TPM2B_AUTH                      66
    TPM2B_DIGEST                    66
    TPMU_SENSITIVE_COMPOSITE        962
Starting ACT thread...
TPM command server listening on port 2321
Platform server listening on port 2322
Command IPv4 client accepted
Platform IPv4 client accepted
^C
/ # 
```

To test type the TPM command `tpm2_getrandom` to obtain a random number:

```bash
tpm2_getrandom 32 --hex
e0b2c1bb096a81032ec1114cc504795ee77c7cc3d159e76165801b38d892296f
/ # 
```

NB: you should get a different random number that we do above!


## Course Material

All the course material is in the [./docs](./docs) directory. Read the `README.md` file first and then work through the documents in order. You can find a set of files that explain different TPM concepts and contain code examples you can use with the container to get familiarized with TPM and its use cases.


## Technologies

Software needed for the actual end functionality is the TCG TPM2 Software Stack
(TSS2) and the IBM TPM simulator.

*  [TPM2-tss](https://github.com/tpm2-software/tpm2-tss) BSD-2-Clause
*  [TPM2-abrmd](https://github.com/tpm2-software/tpm2-abrmd) BSD-2-Clause
*  [TPM2-tools](https://github.com/tpm2-software/tpm2-tools) BSD-3-Clause
*  [IBM simulator](./licenses/LICENSE-ibm-tpm-simulator)



# Credits

Ian Oliver, Nokia Bell Labs
Ronny Backman, Nokia Bell Labs
Gabriela Limonta, Nokia Bell Labs
Jerome Lenssen, Nokia Mobile Networks