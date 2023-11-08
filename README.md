# Nokia TPM Course

This is a short course on getting started with understanding how a TPM 2.0 works. In this course we explain a number of the features of the TPM 2.0 through the TPM2_Tools through examples and, optionally, exercises.

While this course is designed to be used as part of a tutorial session it can be used stand alone and also as an easy reference.

(C)2022 Nokia

## Installation

There are other ways to complete the install, other ways to do each step, but this is common and clear.

Start your install at a terminal window on a linux system.  If you are using a Windows computer, create a linux VM (possibly using VirtualBox), such as the latest Ubuntu system, and begin your install from the linux VM command line.

While in your home directory, install the TPMCourse in your linux VM by using git.  The command is

```bash
git clone https://github.com/nokia/TPMCourse.git
```

Git will clone (copy) files from the remote TPMCourse repository on the Nokia github to your linux system.  However, your linux system may not yet have the git tool.  If you don't have it, you will need to install it in order to complete the git clone. (The linux system should tell you how to install it as part of its negative response to your initial attempt.)  
https://git-scm.com/book/en/v2/Getting-Started-Installing-Git

When the clone is complete, your home directory should have a new TPMCourse sub-directory.  Move to this directory on your linux system and list the files and directories. It should look like the Nokia github TPMCourse repository from which you copied, including an alpine directory, a docs directory, a src directory, a license file and a README file.  Now you are ready to build the container that will host the TPM simulator, the TPM tools, and the alpine linux operating system which is a simple secure linux that will always have the libraries and configurations needed for this course. 

Note, for now, let's assume you are going to work on a TPM Simulator, not an actual TPM in your system hardware.  This course could be adapted to work with a real TPM, but that will not be addressed here.  The TPM simulator is convenient for learning, because if you mess up something, only the simulator is affected and the simulator can easily be restored (by restart). 

To build the TPMCourse container, docker must be available on your linux VM.  If you enter the docker build command below and it starts the 29 step build, you are on your way.  If it fails, the system will provide the instructions on how to install docker on your system.  After docker install, you may receive "Got permission denied while trying to connect to the Docker daemon socket...".  If so, you have a permission issue.  See "How to Fix Docker Permission Denied Error on Ubuntu":
https://linuxhandbook.com/docker-permission-denied/#:~:text=deal%20with%20it.-,Fix%201%3A%20Run%20all%20the%20docker%20commands%20with%20sudo,the%20Docker%20daemon%20socket'%20anymore.


```bash
cd alpine
docker build -t tpmcourse:latest .
```

NOTE: the above might need to be run as sudo, eg:

```bash
sudo docker build -t tpmcourse:latest .
```


To see how the build executes, see the Dockerfile in the alpine directory which when built downloads the various libraries and tools to interact with a TPM and also the IBM TPM Simulator

If all works successfully, after about 29 stages, you can type the following command to run the container in interactive mode.

```bash
docker run -it tpmcourse:latest
```

NOTE: the above might need to be run as sudo, eg:

```bash
sudo docker run -it tpmcourse:latest
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

NOTE: sometimes the `#`` prompt doesn't appear...hit enter, the it should.

To test type the TPM command `tpm2_getrandom` to obtain a random number:

```bash
tpm2_getrandom 32 --hex
e0b2c1bb096a81032ec1114cc504795ee77c7cc3d159e76165801b38d892296f
/ # 
```

NB: you should get a different random number that we do above!

## Caveats

Firstly, you can run everything here against your own, real, hardware (or firmware) TPM - there are some commands which might be irreperable changes - these are clearly noted. We are NOT responsible for any actions resulting in the bricking or worse of your (or someone else's) computer.

The supplied docker file is the best way to explore things in safety...if you break something you can just restart the container :-)

Secondly, sometimes parameters to commands change - we use the latest tpm2_tools and sometimes things do change and the course might not be fully updated. In this case, either make an issue in github or make a fork, change and then a pull request - your contributions either way will be very much appreciated.


## Course Material

All the course material is in the [./docs](./docs) directory. Read the `STARTHERE.md` file first and then work through the documents in order. You can find a set of files that explain different TPM concepts and contain code examples you can use with the container to get familiarized with TPM and its use cases.


## Technologies

Software needed for the actual end functionality is the TCG TPM2 Software Stack
(TSS2) and the IBM TPM simulator.

*  [TPM2-tss](https://github.com/tpm2-software/tpm2-tss) BSD-2-Clause
*  [TPM2-abrmd](https://github.com/tpm2-software/tpm2-abrmd) BSD-2-Clause
*  [TPM2-tools](https://github.com/tpm2-software/tpm2-tools) BSD-3-Clause
*  [IBM simulator](./licenses/LICENSE-ibm-tpm-simulator)


