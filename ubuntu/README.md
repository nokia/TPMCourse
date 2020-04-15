# TPM 2.0 docker container for experimenting

## Licenses

The base image for the container is **ubuntu:disco** (Ubuntu 19.04 / Debian
Buster - Sid).

Dependecies for building the TCG TPM2 Software Stack are:

*  autoconf-archive
*  libcmocka0
*  libcmocka-dev
*  procps
*  iproute2
*  build-essential
*  git
*  pkg-config
*  gcc
*  libtool
*  automake
*  libssl-dev
*  uthash-dev
*  autoconf
*  doxygen
*  libglib2.0-dev
*  libdbus-1-dev
*  libcurl4-gnutls-dev
*  libgcrypt20-dev
*  wget

It is possible to further strip down dependecies, some packages are only used for
building the documentation which is not in use inside the container as
default. Some of the dependecies are also necessary for the software stack to
funcion properly (libcrypt20-dev).

