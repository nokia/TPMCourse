# Quoting

When a machine boots, the state of that machine (hashes of firmware, bootloader, kernel, etc) is stored in the PCRs.

To determine if a machine is trusted, we ask the machine to report its measurements and we compare them to stored reference values.

The process of requesting the measurements from the TPM of a machine is called *quoting*.

We can request a quote for a set of PCRs and the TPM will return a structure called TPMS_ATTEST.
We use the following command to get a quote from the TPM:

```bash
[fedora@localhost ~]$ tpm2_quote -c 0x81010003 -l sha1:0,1,2,3,4,5,6,7,8,9 -m quote.msg -s quote.sig -g sha256 -q 123456 -o pcrs.out
```

This command asks the TPM to produce a quote structure (`TPMS_ATTEST`) which includes a summary of the requested measurements of the platform.
This means that we can request the measurements for a specific set of PCRs (e.g. 1, 2 and 3) and the quote will include a hash over the values in those registers.
This value is a summary of the requested registers because any change in any of the registers will change the output of the hash.
This summary value is called *attested value*.

For example, this is how an attested value for a set of PCRs 0-2 would be built:
```
pcrs:
  sha1:
    0 : 0x7EBA0CFB74F41FEBCCDD1251F02BC208052B6023
    1 : 0xB8720B5234E2F08CFA87069F172CBB43F0F08225
    2 : 0x86FA03B9C721AF57DE8FB1C43CC3FE7B0A42239A
```

```
attested_value = sha256(pcr0 || pcr1 || pcr2)
attested_value = sha256( 7EBA0CFB74F41FEBCCDD1251F02BC208052B6023 || B8720B5234E2F08CFA87069F172CBB43F0F08225 || 86FA03B9C721AF57DE8FB1C43CC3FE7B0A42239A )
attested_value = 54731a1a4664db0dccf8b7c0723f4bf30cd93ffce18c4e4a63e7b8776908007a
```
This can be verified, one way to do it:
```python
Python 2.7.17 (default, Nov  7 2019, 10:07:09) 
[GCC 7.4.0] on linux2
Type "help", "copyright", "credits" or "license" for more information.
>>> import hashlib
>>> hash = hashlib.sha256("7EBA0CFB74F41FEBCCDD1251F02BC208052B6023B8720B5234E2F08CFA87069F172CBB43F0F0822586FA03B9C721AF57DE8FB1C43CC3FE7B0A42239A".decode("hex"))
>>> hash.hexdigest()
'54731a1a4664db0dccf8b7c0723f4bf30cd93ffce18c4e4a63e7b8776908007a'
```

We used the following options for the command.

* `-c 0x81010003` tells the TPM to use this handle (AK) to sign the generated quote.
* `-l sha1:0,1,2,3,4,5,6,7,8,9` tells the TPM to include PCRs 0-9 of the SHA1 bank in the quote.
* `-m quote.msg` tells the tools to store the quote structure in `quote.msg` file.
* `-s quote.sig` tells the tools to store the signature for the quote in the `quote.sig` file.
* `-g sha256` tells the TPM to use the `SHA256` algorithm.
* `-q 123456` tells the TPM to use `123456` as a nonce (to prevent replay attacks.)
* `-o pcrs.out` stores a record of the values of the PCRs indicated with the `-l` option in the file `pcrs.out`.

and this is what the output looks like:

```bash
[fedora@localhost ~]$ tpm2_quote -c 0x81010003 -l sha1:0,1,2,3,4,5,6,7,8,9 -m quote.msg -s quote.sig -g sha256 -q 123456 -o pcrs.out
quoted: ff54434780180022000b36ec8291b370f278c241fe44260da8b24f7bc472879d13c4888017643de294080003123456000000000007b15e000000010000000001201706190016363600000001000403ff03000020900e54b2767b470bf08fb69a1270723a6e2b0f44c661bce7b4a89244a077f9cb
signature:
  alg: rsassa
  sig: 418beea652767313f9d32ec4df7f311015d7d18075d1dd58ef992de4a3a9d1eb3947a34ff8f419cc82df51003e8597e079debcf12debd512d47ac6fe95855e508ea5cbe1f4e32d90a7bf56a4ab561d82108ef8523518f9cd8dd90493815ac1ce3afcfbc1b74f81222f74c4fbf414d423175ded3c5e6db929d77d04d83f282a01e49abce38485e1f92ce847f0a3408e4399d1a1649f7aaea8375e99bbee4fd243216630bfa6afc6d06ba942ed84c59f98bf6f5abdfbf4fb7787a6de32ad23e441e32726916e670745036494466b2c12c2769923c2e99728509a5aa8938e8a2a5327b3714f6f7d82d69e8a499abed6898ac69809ff453224c61dd83714cf5fa55d
pcrs:
  sha1:
    0 : 0x7EBA0CFB74F41FEBCCDD1251F02BC208052B6023
    1 : 0xB8720B5234E2F08CFA87069F172CBB43F0F08225
    2 : 0x86FA03B9C721AF57DE8FB1C43CC3FE7B0A42239A
    3 : 0xB2A83B0EBF2F8374299A5B2BDFC31EA955AD7236
    4 : 0xAB705AAE41789E02A1909B2CBB8BFC0806115004
    5 : 0x7B25D2EABBA18DC910E724A0C75020F1FEC80BE2
    6 : 0xB2A83B0EBF2F8374299A5B2BDFC31EA955AD7236
    7 : 0x518BD167271FBB64589C61E43D8C0165861431D8
    8 : 0xC3DF1A5D37AC51163C63320F28B9C1C1FD933BD2
    9 : 0x944C3EFEB668CB217B260F7D4B594DA4341E6FF2
calcDigest: 900e54b2767b470bf08fb69a1270723a6e2b0f44c661bce7b4a89244a077f9cb
```

The quote structure contains other interesting fields apart from the `attested` value.
It includes:
* The nonce used with option `-q` (if used)
* A clock value
* A reset counter (number of times the machine has been rebooted)
* A restart counter (number of times the machine has been suspended or hibernated)
* A safe value which indicates whether the clock value is valid
* Firmware version of the TPM
* Information about the key used to sign the quote

We can decode those fields with the following command:
```bash
[fedora@localhost ~]$ tpm2_print -t TPMS_ATTEST quote.msg
magic: ff544347
type: 8018
qualifiedSigner: 000b36ec8291b370f278c241fe44260da8b24f7bc472879d13c4888017643de29408
extraData: 123456
clockInfo:
  clock: 504158
  resetCount: 1
  restartCount: 0
  safe: 1
firmwareVersion: 2017061900163636
attested:
  quote:
    pcrSelect:
      count: 1
      pcrSelections:
        0:
          hash: 4 (sha1)
          sizeofSelect: 3
          pcrSelect: ff0300
    pcrDigest: 900e54b2767b470bf08fb69a1270723a6e2b0f44c661bce7b4a89244a077f9cb
```
