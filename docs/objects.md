# TPM Objects
- [TPM Objects](#tpm-objects)
  - [Handles](#handles)
  - [Hierarchies](#hierarchies)
    - [Taking Ownership](#taking-ownership)
    - [Dictionary Lockout](#dictionary-lockout)
    - [Clearing the TPM](#clearing-the-tpm)
    - [Resetting Seeds](#resetting-seeds)
  - [Attributes](#attributes)
  - [Locality](#locality)

The TPM contains a number of objects such as keys, NVRAM areas, PCRs etc. These are addressed using handles - basically a pointer to the object.

Each object exists in one of the four authorisation hierarchies in the TPM: platform, owner, endorsement and null - these will be explained later.

Each object has a set of attributes associated with it that provides information about what kinds of authorisation are required, information about the object itself and so on.

Finally, there is a notion called locality which further controls who can access and modify objects.

## Handles
A handle is a 32 bit number that is used to refer to an object stored in the TPM. Typically we write these in hexadecimal notation.

The upper byte(s) tells us what kind of object the handle points to.

   * `0x80......` - transient objects
   * `0x8100....` - owner hierarchy keys
   * `0x8101....` - endorsement hierarchy keys
   * `0x8180....` - platofrm hierachy keys
   * `0x01......` - NVRAM
   * `0x01c1....` - NVRAM areas that the TCG has reserved for particular organisations/uses
   * `0x01c2....` - NVRAM areas that the TCG has reserved for OEMs, eg: Infineon
   * `0x01c2....` - NVRAM areas that the TCG has reserved for platform OEMs     
   * `0x40......` - Permanent objects (eg: references to hierarchies)

The handles can be listed using the `tpm2_getcap` command.

   * `-l` lists the capability categories of the TPM, for example:

```bash
$ tpm2_getcap -l
- algorithms
- commands
- pcrs
- properties-fixed
- properties-variable
- ecc-curves
- handles-transient
- handles-persistent
- handles-permanent
- handles-pcr
- handles-nv-index
- handles-loaded-session
- handles-saved-session
$ tpm2_getcap handles-permanent
- 0x40000001
- 0x40000007
- 0x40000009
- 0x4000000A
- 0x4000000B
- 0x4000000C
- 0x4000000D
```

These handles refer to the various authorisation hierarchies: platform, owner, endorsement and null; the lockout password and other structures inside the TPM.


## Hierarchies
The TPM has four hierachies that are used to group objects together for specific use cases. The hierarchies are:

   * Platform
   * Owner (sometimes called Storage)
   * Endorsement
   * Null

Each hierarchy is used to define an authorisation to access and use the objects it contains. Typically this is done using a password.

One use of the hierarchies is during the manufacturing and provisioning of a TPM:

   * The manufacturer may place keys or NVRAM areas under the Endorsement hierarchy and then lock it
   * The OEM may place keys or NVRAM areas under the Platform hierarchy and then lock it (see note below)
   * The Owner hierarchy is open to the end-user
   * The null hierarchy is used internally by the TPM for anything temporary, eg: session information etc. 

The reason for locking these areas is to facilitate revocation of keys and other data, as well as keeping manufacturer or OEM specific data safe.

Each hierarchy has a seed from which keys are derived. It is possible - on some TPMs - to reset these seeds rendering any keys, policies and NVRAM areas under that hierarhy no longer accessible or usable.

The platform hierarchy is interesting as its password is always reset when a power cycle happens. This is because this hierarchy was aimed at platform manufacturers and OEMs who want control over their parts of the system, eg: firmware signatures etc. But, if something needed changing, eg: a firmware update, then it would be necessary to unlock this hierarchy. The solution is that the platform hierarchy is open until device start and then a random password written there to lock it for the duration of usage. Unlocking this hierarchy would be part of a hardware update process. 

In these exercises use the owner hierarchy - it is what it was designed for.

### Taking Ownership
Note, there used to be a command called `takeownership` which is kind of a hangover from the TPM 1.2 days when such a concept existed. It was modified to call the TPM 2.0 equivalents and has over the years caused some confusion about what taking ownership actually meant. TPM 2.0 has hierarchies, while TPM 1.2 didn't ... if someone is talking about taking ownership then it is time for a longer conversation on exactly what they want. This is not a discussion we will have here but it is good to be aware.

To set the passwords on the hierarchies, typically when a TPM is first used, we set three passwords using `tpm2_changeauth`.

   * `owner` or `o` sets the owner hierarchy password
   * `endorsement` or `e` sets the endorsement hierarchy password
   * `lockout` or `l` sets the lockout password

WARNING #1: This is a potentially dangerous command. Practice with the simulator and not on a real TPM. There will be more warnings later.

```bash
$ tpm2_changeauth -c owner password1 
$ tpm2_changeauth -c endorsement password2
$ tpm2_changeauth -c lockout password3
```

The lockout password is used the unlock the password from the dictionary attack state. If someone enters a password incorrectly mulitple times the TPM will lock. In some cases this lockout period might be over 24 hours!

WARNING #2: We recommend when testing and practicing the commands here that you DO NOT SET ANY PASSWORDS.

WARNING #3: If you forget or incorrectly set the passwords then you will lose access to the TPM....forever!

To change the password, supply the old password via `-p`  and the new password.

```bash
$ tpm2_changeauth -c o -p password1 passwordA
$ tpm2_changeauth -c e -p password2 passwordB
$ tpm2_changeauth -c l -p password3 passwordC
```

You can also clear all the passwords by supplying the lockout password like so




### Dictionary Lockout
WARNING: THIS COULD RESULT IN YOU BEING LOCKED OUT OF YOUR TPM FOR EITHER A LONG TIME OR PERMANENTLY!

To prevent continuous attempt to access objects on a TPM, the device employs a protection called dictionary lockout.

When the TPM fails with a lockout error then it is necesary to supply the dictionary lockout password - stored one of the permanent handles and set with the `tpm2_changeauth` command.

The command for setting the properites of the lockout is as follows:

```bash
tpm2_dictionarylockout -s -n 5 -t 6 -l 7 -p passwordC
```

TODO: rewrite this so it isn't a copy of the man page.

   * `-s` Specifies the tool should operate to setup dictionary-attack-lockout parameters.
   * `-c` clear-lockout: Specifies the tool should operate to clear dictionary-attack-lockout state.
   * `-l` --lockout-recovery-time=NATURAL_NUMBER. Specifies the wait time in seconds before another TPM_RH_LOCKOUT authentication attempt can be made after a failed authentication.
   * `-t` --recovery-time=NATURAL_NUMBER: Specifies the wait time in seconds before another DA-protected-object authentication attempt can be made after max-tries number of failed authentications.
   * `-n` --max-tries=NATURAL_NUMBER: Specifies the maximum number of allowed authentication attempts on DA-protected-object; after which DA is activated.
   * `-p` --auth=AUTH: The authorization value for the lockout handle.




### Clearing the TPM
WARNING: THIS IS A VERY DANGEROUS, NON-REVERSABLE OPERATION WHICH WILL RESET THE TPM AND REGENERATE THE SEEDS.

To clear a hierarchy, we need to specify it with the `-c` option.
We can use  `p` for plaform, `e` for endorsement and `o` for owner.
For example, to clear the platform hierarchy, we would use the following command:

```bash
$ tpm2_clear -c p
```

We can also clear the lockout password mentioned in the previous section, by using `l` as an argument for the `-c` option:

```bash
$ tpm2_clear -c l
```




### Resetting Seeds
WARNING: THIS IS A VERY DANGEROUS, NON-REVERSABLE OPERATION WHICH WILL RESET THE TPM AND REGENERATE THE SEEDS.

To reset the seeds to generate the primary keys for the platform and endorsement hierarchies:

```bash
$ tpm2_changepps
```

If a password has been set with `tpm2_changeauth` then the above commands take the option `-p` followed by the password for that hierarchy.

```bash
$ tpm2_changepps -p <password>
```


## Locality
Locality is a mechanism that is used to tell the TPM what state the system is in.
When locality is applied to PCR registers, it controls who can extend and reset those registers.

The TCG defines a number of localities, of which 0 to 4 are typically in use.
To use other localities one must refer to the TCG.
Examples of usage are below:

   * Locality 0 is used to control access to PCRs that can be modified and reset by the CRTM
   * Locality 1 is used to control access to PCRs that can be modified and reset by the SRTM
   * Locality 2 is used to control access to PCRs that can be modified and reset by the Late launch
   * Locality 3 is used to control access to PCRs that can be modified and reset by the Kernel
   * Locality 4 is used to control access to PCRs that can be modified and reset by the User/Run-time environemnt

For example, in normal operation PCRs 16 and 23 are defined under locality 4 which means the user can reset and extend these PCRs.
PCR 0 is under locality 0, which means only the intial CRTM code can modify this PCR.

Locality is a complex topic but knowing that the TPM might refuse certain operations on PCRs and this is the reason why will help in understanding some errors. Locality is described earlier in this tutorial and normally we don't worry about it.

Locality is a complex topic but knowing that the TPM might refuse certain operations on PCRs and this is the reason why will help in understanding some errors. Locality is described earlier in this tutorial and normally we don't worry about it.

