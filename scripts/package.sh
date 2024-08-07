#!/bin/bash
set -e

if [ -n "$1" ]; then
    DEST=$1
else
    DEST="${PWD}/EloqDM"
fi
mkdir -p ${DEST}/bin
mkdir -p ${DEST}/logs
mkdir -p ${DEST}/config
cp dm/config-example/* ${DEST}/config/

make dm
cp bin/dm* ${DEST}/bin/
pushd ../tidb
make build_dumpling
cp bin/dumpling ${DEST}/bin/
make build_lightning
cp bin/tidb-lightning ${DEST}/bin/lightning
make build_lightning-ctl
cp bin/tidb-lightning-ctl ${DEST}/bin/lightning-ctl
popd
tar -czvf eloqdm.tar.gz EloqDM
echo "Done!"
