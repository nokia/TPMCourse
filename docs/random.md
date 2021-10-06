# Random Numbers
- [Random Numbers](#random-numbers)
  - [Random Number in Raw Format](#random-number-in-raw-format)
  - [Random Number is Hexadecimal](#random-number-is-hexadecimal)
  - [Additional Randomness](#additional-randomness)

The simplest command to use with a TPM is one to obtain random numbers from the onboard hardware random number generator.

## Random Number in Raw Format
The command `tpm2_getrandom` accepts an integer as input and returns that amount of random bytes.

The minimum number of bytes to return is 1

The maximum number of bytes to return may vary. Typically 32 or 64 might be the upper limit - this depends upon how the manufacturer has built their TPM.

```bash
$ tpm2_getrandom 20
�
 h�HU��ns�4�y�'
```

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

# Exercises

1.  Generate a single random value 
2.  What is the larger amount of random values your TPM can generate+
3.  Write a small script that outputs a random value between, say 1 and 6.  a TPM backed dice.