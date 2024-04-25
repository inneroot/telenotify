#!/bin/bash

echo "readint pg.env"
source pg.env

if [ "$MIGRATION_CONN_STRING" = "" ]; then
    echo "MIGRATION_CONN_STRING is not set, no migrations"
    exit 0
fi

echo "tern migrate"
tern migrate --migrations ./migrations/ --conn-string "$MIGRATION_CONN_STRING"
