#!/bin/bash
set -eu

# For security measures, daemons should not be run as sudo. Execute juju-dns as the non-sudo user: snap_daemon.
export LOCPATH="${SNAP}"/usr/lib/locale
exec "${SNAP}"/bin/juju-dns -conf "${SNAP_DATA}"/etc/Corefile "$@"
