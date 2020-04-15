# TPM 2.0 alpine container

The base image for the container is **alpine:latest**.

Dependecies for building the TCG TPM2 Software Stack are:

* autoconf
* autoconf-archive
* automake
* libtool
* build-base
* pkgconf
* doxygen
* json-c-dev
* openssl
* openssl-dev
* libssl1.1
* git
* udev 
* dbus
* curl-dev
* linux-headers
* glib-dev
* libconfig-dev
* libgcrypt-dev
* wget

It is possible to further strip down dependecies, some packages are only used for
building the documentation which is not in use inside the container as
default.

