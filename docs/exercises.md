# Exercises

Some interesting exercises for you - this is incomplete but, let's see...

## Random Numbers

3 Simple exercises. Number #3 requires some bash/csh script programming.

1.  Generate a single random value 
2.  What is the largest amount of random values your TPM can generate+
3.  Write a small script that outputs a random value between, say 1 and 6.  a TPM backed dice.

## Objects

1. What are the properties of your TPM?
2. Using `tpm2_getcap` find which PCR banks your TPM supports
3. What permanent objects are stored in the TPM? What are they exactly?
4. What cryptographic algorithms does your TPM support?


## PCRs

1. What do the PCRs on your machine report?
2. What do each of the PCRs mean?   x86 UEFI machines and Linux/Windows, refer to the TCG documentation on this. Actually TianoCore's documentation has a good description of measured boot and what the PCRs mean - go find this.
3. Do any change after reboot?
4. Try entering the UEFI setup page and then letting the system boot...anything?

## Keys

1. Generate the PEM representations for the EK and AK - how would these be use to identify the TPM?
2. If you have access to two (or more) TPMs, generate your own keys on each, distribute the public keys and try sending encrypted messages between them, ie: encrypt on one TPM, decrypt on another...
