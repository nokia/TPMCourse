#!/bin/bash

# Start dbus service and sleep for 1 second to wait that it has started.
# If tpm2-abrmd starts before the dbus service is ready we get an error.
service dbus start
sleep 1

# Start the ibm tpm simulator with default settings
/tpm2/ibmsim/src/tpm_server &

# Start tpm2 access broker & resource manager daemon
tpm2-abrmd --allow-root --tcti=mssim &

# Continue with a shell as default
bash

# Future development to start as user and not root
# dbus-daemon --config-file=/usr/share/dbus-1/session.conf --fork
