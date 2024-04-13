#!/bin/bash
source migration.env

# set -e
if [ "$MIGRATION_CONN_STRING" = "" ]; then
    echo "MIGRATION_CONN_STRING is not set, no migrations"
    exit 0
fi
tern migrate --migrations ./migrations/ --conn-string "$MIGRATION_CONN_STRING"
