#!/bin/sh

set -e

echo "db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up


echo "start the app"
exec "$@"