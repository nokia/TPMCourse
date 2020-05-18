# Getting started with TPM

Here you can find some documentation on TPM functionality and different commands you can run.
These files are meant to introduce basic concepts and commands that you can use on the TPM.
The code examples included in these files can be tested with the simulator provided by the docker container available in this repository.

A brief summary of the contents:

* [Command Summary](./commandSummary.md) provides a summary of some of the TPM commands available and their posible use cases.
* [Objects](./objects.md) introduces the different kinds of objects you can store in a TPM and how to access and use them.
* [Keys](./keys.md) contains an explanation of how the keys in the TPM work, as well as how to generate, load and use keys in the TPM.
* [NVRAM](./nvram.md) explains how to use the TPM's Non-Volatile RAM as secure storage.
* [PCRs](./pcrs.md) introduces the Platform configuration Registers (PCRs), which are used for storing measurements of a platform.
* [Quoting](./quoting.md) covers the quoting mechanism, which is used to obtain measuremens from a TPM for remote attestation purposes.
* [Random](./random.md) covers the random number generator capabilities on the TPM.