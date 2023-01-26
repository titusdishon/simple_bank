#!/bin/sh

set -e # script exits immediately if command returns non-zero status 

echo "run db migrations"
source /app/app.env

/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"

exec "$@"
