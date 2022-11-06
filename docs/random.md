# Random Numbers
- [Random Numbers](#random-numbers)
  - [Random Number in Raw Format](#random-number-in-raw-format)
  - [Random Number is Hexadecimal](#random-number-is-hexadecimal)
  - [Additional Randomness](#additional-randomness)

The simplest command to use with a TPM is one to obtain random numbers from the onboard hardware random number generator. This command is a useful one as you can always use it to test if your TPM is working.

The TPM can return a limited number of bytes at a time, typically between 1 and 32, though this might vary by manufacturer.

## Random Number in Raw Format
The command `tpm2_getrandom` accepts an integer as input and returns that amount of random bytes.

The minimum number of bytes to return is 1

The maximum number of bytes to return may vary. Typically 32 or 64 might be the upper limit - this depends upon how the manufacturer has built their TPM.

```bash
$ tpm2_getrandom 20
�
 h�HU��ns�4�y�'
```

Note: the value returned above *is* random so what you get back should be different - the more you return the greater the probability.

## Random Number is Hexadecimal
To return the bytes in hexadecimal format include the option `--hex`

```bash
$ tpm2_getrandom 20 --hex
01654e15b23415b029c8cdc0085ac7f146c5a1c6
```


## Additional Randomness   
The TPM has an operation to add more entropy to the random number generator.

The command `tpm2_stirrandom` takes up to 128 bytes from a file or piped in from stdin to achieve this. Some examples below of two mechanisms:

```bash
$ echo -n "myrandomdata" | tpm2_stirrandom
$ dd if=/dev/urandom bs=1 count=64 > myrandom.bin
$ tpm2_stirrandom < ./myrandom.bin
```
