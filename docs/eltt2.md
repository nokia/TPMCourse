#ELTT2

Infineon make a small opensource tool for interacting with the TPM 2 device called $eltt2$. While not as sophisticated as the tpm2tools or the functionality provided by, say, the GO-TPM or WolfTPM libreries, it is very small, simple and can run on small embedded devices

Eltt2 is primarily designed as a debugging tool, it should not be relied upon to provide attestation data etc. However it does run on practically anything (at least in our experience) :-)

## Help

```bash
$eltt2 -h

'-a <data bytes>': Hash Sequence SHA-1
        -> Data bytes: Enter a byte sequence like '0F56...' for {0x0f, 0x56, ...}
'-A <data bytes>': Hash Sequence SHA-256
        -> Data bytes: Enter a byte sequence like '0F56...' for {0x0f, 0x56, ...}
'-b <command bytes>': Enter your own TPM command
        -> Command bytes: Enter your command bytes in hex like '0f56...' for {0x0f, 0x56, ...}
'-c': Read Clock
'-d <shutdown type>': Shutdown
        -> Shutdown types: clear [default], state
'-e <PCR index> <PCR digest>': PCR Extend SHA-1
        -> PCR index:  Enter the PCR index in hex like '17' for 0x17
           PCR digest: Enter the value to extend the PCR with in hex like '0f56...' for {0x0f, 0x56, ...}
'-E <PCR index> <PCR digest>': PCR Extend SHA-256
        -> PCR index:  Enter the PCR index in hex like '17' for 0x17
           PCR digest: Enter the value to extend the PCR with in hex like '0f56...' for {0x0f, 0x56, ...}
'-g': Get fixed capability values
'-v': Get variable capability values
'-G <byte count>': Get Random
        -> Enter desired number of random bytes in hex like '20' for 0x20 (=32 bytes, maximum)
'-h': Help
'-r <PCR index>': PCR Read SHA-1
        -> PCR index: Enter PCR number in hex like '17' for 0x17
'-R <PCR index>': PCR Read SHA-256
        -> PCR index: Enter PCR number in hex like '17' for 0x17
'-s <data bytes>': Hash SHA-1
        -> Data bytes: Enter a byte sequence like '0F56...' for {0x0f, 0x56, ...}
'-S <data bytes>': Hash SHA-256
        -> Data bytes: Enter a byte sequence like '0F56...' for {0x0f, 0x56, ...}
'-t <selftest type>': SelfTest
        -> Selftest type: not_full [default], full, incremental
'-T': Get Test Result
'-u <startup type>': Startup
        -> Startup types: clear [default], state
'-z <PCR index>': PCR Reset SHA-1 and SHA-256
        -> PCR index: Enter PCR number in hex like '17' for 0x17
```

Note, the examples here were run with `sudo` - this was because of the machine not being set up to allow access to the `/dev/tpm0` device by normal users. Typically the device is owned by the user and group `tss` and adding a normal user to this group will allow access without sudo. 

## Random Numbres

$ sudo ./eltt2 -G 10
Random value:

0x00000000:   0xA5  0x36  0xB6  0xBF  0xA5  0x8B  0xD0  0x59  0x90  0xFD  0x1F  0x08  0x37  0x6A  0xF3  0xC9  



## Hashing
Eltt2 can hash suitablly presented data. The `-a` and `-s` options give access to the SHA-1 functions, while `-A` and `-S` to the SHA-256 functions. Eltt2 as you might notice is verbose, but then again it is more of a debugging and testing tool:

```bash
~/Work/eltt2$ sudo ./eltt2 -a abc


TPM2_HashSequenceStart of 'abc' with SHA-1:
TPM Response:
80 01                         TPM TAG
00 00 00 0E                   RESPONSE SIZE
00 00 00 00                   RETURN CODE
 Command-specific response Data:
80 00 00 00 

TPM2_SequenceUpdate:
TPM Response:
80 02                         TPM TAG
00 00 00 13                   RESPONSE SIZE
00 00 00 00                   RETURN CODE
 Command-specific response Data:
00 00 00 00 00 00 01 00 00 

TPM2_SequenceComplete:
TPM Response:
80 02                         TPM TAG
00 00 00 31                   RESPONSE SIZE
00 00 00 00                   RETURN CODE
 Command-specific response Data:
00 00 00 1E 00 14 DC AA 4E 54 
32 32 BB B6 FD CB 68 D8 AC 4C 
55 73 F3 6E 10 CA 80 24 40 00 
00 07 00 00 00 00 01 00 00 

Hash value extracted from TPM response:

0x00000000:   DC  AA  4E  54  32  32  BB  B6  
0x00000008:   FD  CB  68  D8  AC  4C  55  73  
0x00000010:   F3  6E  10  CA  


~/Work/eltt2$ sudo ./eltt2 -s abc


TPM2_Hash of 'abc' with SHA-1:

0x00000000:   DC  AA  4E  54  32  32  BB  B6  
0x00000008:   FD  CB  68  D8  AC  4C  55  73  
0x00000010:   F3  6E  10  CA  

```


## Clock, Reset, Restart and Safe
Eltt2 can read the TPM 2.0 clock status and provide information about the state of the device. For example, below we can see the uptime of the device, the time since the last TPM clear, the number of reboots (reset), the number of shutdowns - which actually refers to the number of *clean* shutdowns and whether the TPM was resarted from a properly shutdown state (safe):

```bash
~/Work/eltt2$ sudo ./eltt2 -c


Clock info:
=========================================================
Time since the last TPM_Init:
1934920 ms  =  0 y, 0 d, 0 h, 32 min, 14 s, 920 ms

Time during which the TPM has been powered:
35001929 ms  =  0 y, 0 d, 9 h, 43 min, 21 s, 929 ms

TPM Reset since the last TPM2_Clear:            12
Number of times that TPM2_Shutdown:             1
Safe:                                           1 = Yes

```

This same information is provided as part of the TPMS_ATTEST structure when quoting the TPM - see the tpm2tools tpm2_quote command.  It can be used in the attestation process but be aware that the above information is provided as-is by the TPM and is not signed by the TPM.


## Reading and Writing the PCRs
To read a PCR provide the PCR register in hexadecimal. The tool can only access the standard mandated SHA-1 and SHA-256 banks depending upon the commmand line option provided. The tool can only return a single PCR at a time so to obtain multiple PCR values, multiple individual calls must be made.

### Reading

Accessing PCR 0 in the SHA-1 and SHA-256 banks is achieved respectively:

```bash
~/Work/eltt2$ sudo ./eltt2 -r 0

Read PCR 0 (SHA-1):
20 E5 BB F9 0D 46 6A 79 75 7D 63 D9 94 B2 57 C9 02 F2 5F 9E 

~/Work/eltt2$ sudo ./eltt2 -R 0

Read PCR 0 (SHA-256):
79 86 0D C6 51 CE 87 C8 E4 F6 31 C5 FC ED D7 7B 0A 38 20 6A DE BC 14 4A 1B 8A 76 23 59 3E AA E2 
```

Remeber hexadecimal: PCRs 10 to 23 ... dec 10 is hex A and dev 23 is hex 17, note the delberate errors below:

```bash
~/Work/eltt2$ sudo ./eltt2 -r a

Read PCR 10 (SHA-1):
33 D1 BA C6 D2 AF 75 DC 95 55 B7 4E 0C 65 B1 35 27 FF D7 CE 

~/Work/eltt2$ sudo ./eltt2 -r 17

Read PCR 23 (SHA-1):
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 

~/Work/eltt2$ sudo ./eltt2 -r 18

Bad option. Enter a hex value between 0x00 and 0x17 in hex without leading '0x'.
Unexpected error: 0xFFFFFFFE

~/Work/eltt2$ sudo ./eltt2 -r 23

Bad option. Enter a hex value between 0x00 and 0x17 in hex without leading '0x'.
Unexpected error: 0xFFFFFFFE
```

### Extending

Extending PCRs is similar straightforward with the correct size of digest being 0-padded. For example, extending PCR 23 )hex 17) with a digest starting with the bytes hex ABC...

```bash
~/Work/eltt2$ sudo ./eltt2 -e 17 abc

Extend PCR 23 (SHA-1) with digest { AB 0C 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 }:
TPM Response:
80 02                         TPM TAG
00 00 00 13                   RESPONSE SIZE
00 00 00 00                   RETURN CODE
 Command-specific response Data:
00 00 00 00 00 00 01 00 00 

~/Work/eltt2$ sudo ./eltt2 -E 17 abc

Extend PCR 23 (SHA-256) with digest { AB 0C 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 }:
TPM Response:
80 02                         TPM TAG
00 00 00 13                   RESPONSE SIZE
00 00 00 00                   RETURN CODE
 Command-specific response Data:
00 00 00 00 00 00 01 00 00 
```


### Resetting

Once the operating system is up and running and (at least on x86 devices) the TPM is running in locality 4. This means that only PCRs 16 and 23 (hex 10 and 17) are resettable (even when running as root!).

We can reset PCR 23 with the `-z` option as demonstrated below:

```bash
~/Work/eltt2$ sudo ./eltt2 -r 17 

Read PCR 23 (SHA-1):
12 0B 50 96 72 A0 5D C9 2A AE F6 D9 91 73 56 22 FD 5C 95 2A 

~/Work/eltt2$ sudo ./eltt2 -z 17 

Reset PCR 23 (SHA-1 and SHA-256):
Done.


~/Work/eltt2$ sudo ./eltt2 -r 17 

Read PCR 23 (SHA-1):
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 
```


Note: these examples are running in userland on Linux on an x86 machine. If you can run eltt at a different level, eg: as a firmware component, then you'll have much more opportunities to reset whatever you want. On other architectures and platforms, eg: embedded devices, ARM, RISC-V etc, the TPM may be in a different locality due to the design of that system.



## TPM Properties
Finally the tool can prpvide information about the TPM itself

```bash
~/Work/eltt2$ sudo ./eltt2 -g


TPM capability information of fixed properties:
=========================================================
TPM_PT_FAMILY_INDICATOR:        2.0
TPM_PT_LEVEL:                   0
TPM_PT_REVISION:                116
TPM_PT_DAY_OF_YEAR:             15
TPM_PT_YEAR:                    2016
TPM_PT_MANUFACTURER:            STM 
TPM_PT_VENDOR_STRING:           
TPM_PT_VENDOR_TPM_TYPE:         1
TPM_PT_FIRMWARE_VERSION:        71.12.40976.4

TPM_PT_MEMORY:
=========================================================
Shared RAM:                     0 CLEAR
Shared NV:                      0 CLEAR
Object Copied To Ram:           0 CLEAR

TPM_PT_PERMANENT:
=========================================================
Owner Auth Set:                 0 CLEAR
Sendorsement Auth Set:          0 CLEAR
Lockout Auth Set:               0 CLEAR
Disable Clear:                  0 CLEAR
In Lockout:                     0 CLEAR
TPM Generated EPS:              0 CLEAR


~/Work/eltt2$ sudo ./eltt2 -v


TPM capability information of variable properties:

TPM_PT_STARTUP_CLEAR:
=========================================================
Ph Enable:                      1 SET
Sh Enable:                      1 SET
Eh Enable:                      1 SET
Orderly:                        1 SET

```


