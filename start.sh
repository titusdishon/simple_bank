#!/bin/sh

set -e # script exits immediately if command returns non-zero status 

echo "run db migrations"

/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"

exec "$@"
