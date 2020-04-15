#TPM2 simulator
FROM ubuntu:focal
LABEL Maintainer="Ronny Backman <ronny.backman@nokia.com>"
LABEL version="1.0"
LABEL description="Base image for running TPM2 Software Stack (TSS), Access \
Broker & Resource Management Daemon (ABRMD) and tools with IBM tpm simulator"

ENV LC_ALL C.UTF-8
ENV LANG C.UTF-8
ENV LANGUAGE C.UTF-8
ENV DEBIAN_FRONTEND noninteractive

COPY open-source-tpm-ubuntu-spdx.rdf /

# Installing needed dependecies and tools for building

RUN apt update && \
    apt install -y \
    autoconf-archive \
    libcmocka0 \
    libcmocka-dev \
    procps \
    iproute2 \
    build-essential \
    git \
    pkg-config \
    gcc \
    libtool \
    automake \
    libssl-dev \
    uthash-dev \
    autoconf \
    doxygen \
    libjson-c-dev \
    libini-config-dev \
    libglib2.0-dev \
    libdbus-1-dev \
    libcurl4-gnutls-dev \
    libgcrypt20-dev \
    wget \
    && apt autoremove -y


# Cloning, compiling and installing
# tpm2-tss, tpm2-abrmd and tpm2-tools
# Source: https://github.com/tpm2-software

RUN mkdir -p /tpm2
WORKDIR /tpm2
RUN git clone https://github.com/tpm2-software/tpm2-tss.git && \
    git clone https://github.com/tpm2-software/tpm2-abrmd.git && \
    git clone https://github.com/tpm2-software/tpm2-tools.git

WORKDIR /tpm2/tpm2-tss
RUN ./bootstrap \
    && ./configure --enable-unit \
    && make -j$(proc) check \
    && make install \
    && ldconfig \
    && make clean


WORKDIR /tpm2/tpm2-abrmd
RUN ./bootstrap \
    && ./configure --with-dbuspolicydir=/etc/dbus-1/system.d \
       --with-systemdpresetdisable \
       --libdir=/usr/lib \
    && make -j$(nproc) \
    && make install \
    && ldconfig \
    && make clean


WORKDIR /tpm2/tpm2-tools
RUN ./bootstrap \
    && ./configure --prefix=/usr \
    && make -j$(nproc) \
    && make install \
    && make clean


# Installing and configuring the IBM tpm2 simulator
# Source: https://sourceforge.net/projects/ibmswtpm2/
# no-check-certificate can be removed when building outside Nokia intra.

RUN mkdir /tpm2/ibmsim
WORKDIR /tpm2/ibmsim
RUN wget --quiet --show-progress --progress=dot:giga --no-check-certificate "https://downloads.sourceforge.net/project/ibmswtpm2/ibmtpm1563.tar.gz" \
    && tar -xf ibmtpm1563.tar.gz \
    && rm ibmtpm1563.tar.gz


WORKDIR /tpm2/ibmsim/src
RUN make


# Adding startup script and using that as COMMAND on container start
COPY tpmStartup.sh /usr/local/bin
WORKDIR /
ENTRYPOINT ["/bin/bash"]
CMD ["tpmStartup.sh"]
