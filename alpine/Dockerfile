#TPM2 simulator with alpine base
FROM alpine:3.13.6
LABEL Maintainer="Ronny Backman <ronny.backman@nokia.com>"
LABEL version="1.0"
LABEL description="Base image for running TPM2 Software Stack (TSS), Access \
Broker & Resource Management Daemon (ABRMD) and tools against IBM tpm simulator"

ENV LC_ALL C.UTF-8
ENV LANG C.UTF-8
ENV LANGUAGE C.UTF-8

COPY open-source-tpm-alpine-spdx.rdf /

# Installing needed dependecies and tools for developing

RUN apk update && \
    apk add \
    acl \
    autoconf \
    autoconf-archive \
    automake \
    libtool \
    build-base \
    pkgconf \
    doxygen \
    json-c-dev \
    openssl \
    openssl-dev \
    libssl1.1 \
    git \
    udev \
    dbus \
    curl-dev \
    linux-headers \
    glib-dev \
    libconfig-dev \
    libgcrypt-dev \
    shadow \
    wget

RUN mkdir tpm2
WORKDIR /tpm2
RUN mkdir ibmsim

WORKDIR /tpm2
RUN git clone --branch 3.0.x https://github.com/tpm2-software/tpm2-tss.git && \
    git clone https://github.com/tpm2-software/tpm2-abrmd.git && \
    git clone https://github.com/tpm2-software/tpm2-tools.git

WORKDIR /tpm2/tpm2-tss
RUN ./bootstrap && \
    ./configure --with-udevrulesprefix && \
    make -j4 && \
    make install

WORKDIR /tpm2/tpm2-abrmd
RUN ./bootstrap && \
    ./configure --with-dbuspolicydir=/etc/dbus-1/system.d && \
    make -j4 && \
    make install

WORKDIR /tpm2/tpm2-tools
RUN ./bootstrap && \
    ./configure && \
    make -j4 && \
    make install


WORKDIR /tpm2/ibmsim
RUN wget --quiet --show-progress --progress=dot:giga --no-check-certificate "https://downloads.sourceforge.net/project/ibmswtpm2/ibmtpm1661.tar.gz"
RUN tar -xf ibmtpm1661.tar.gz
WORKDIR /tpm2/ibmsim/src
RUN make -j4
RUN cp tpm_server /bin/


WORKDIR /

COPY tpmStartup.sh /bin/

CMD ["tpmStartup.sh"]
