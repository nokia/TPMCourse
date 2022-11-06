# Getting started with TPM

Here you can find tutorials, examples and documentation on TPM functionality and different commands you can run.
These documents are meant to introduce basic concepts and commands that you can use on the TPM.
The code examples included in these files can be tested with the simulator provided by the docker container available in this repository or you can test these against a real TPM.

The whole course contains information relating to some of the major tool sets and libraries associated with the TPM. Currently are described basic usages for the tpm2tools command line tools for Linux and Windows, though you might notice some other files, eg: for ELTT2 which is a very simple tool from Infineon.

## Contents 

While there is no strict order, if this is the first time you are here or you wish to work through in a proper methodological way then this is the order you should work through this course.

* [Random](./random.md) covers the random number generator capabilities on the TPM.
* [Objects](./objects.md) introduces the different kinds of objects you can store in a TPM and how to access and use them.
* [Keys](./keys.md) contains an explanation of how the keys in the TPM work, as well as how to generate, load and use keys in the TPM.
* [NVRAM](./nvram.md) explains how to use the TPM's Non-Volatile RAM as secure storage.
* [PCRs](./pcrs.md) introduces the Platform configuration Registers (PCRs), which are used for storing measurements of a platform.
* [Quoting](./quoting.md) covers the quoting mechanism, which is used to obtain measuremens from a TPM for remote attestation purposes.

A Summary of all commands introduced is found in:

* [Command Summary](./commandSummary.md) provides a summary of some of the TPM commands available and their posible use cases.

A set of exercises for the above can be found here:

* [Exercises](./exercises.md) exercises (surprisingly)


## ELTT2

While the eltt2 code is not included in our container we have provided a short introduction to its operation. 

* [ELTT](./eltt2.md) covers basic usage of Infineon's eltt2 tool for embedded systems. Very useful for all devices!
