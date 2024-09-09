#!/bin/sh
set -e

echo "Starting backend..."

ls -a
cd backend

exec "$@"