#!/bin/sh
set -e

cd /docsystem/

if [ ! -f "/docsystem/conf/app.conf" ] ; then
    cp /docsystem/conf/app.conf.example /docsystem/conf/app.conf
fi

export ZONEINFO=/docsystem/lib/time/zoneinfo.zip
/docsystem/docsystem_linux_amd64 install
/docsystem/docsystem_linux_amd64