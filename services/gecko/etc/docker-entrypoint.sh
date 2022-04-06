#!/usr/bin/with-contenv bash

set -e # Exit build script on first failure.
set -u # Exit on unset variable.
set -x # Echo commands to stdout.

# ls -la /app/

# if [ -f "$DB_PATH" ]; then
#   echo "Database already exists, skipping restore"
# else
#   echo "No database found, restoring from replica if exists"
#   res=$(
#     exec /bin/litestream restore -v -if-replica-exists "${DB_PATH}"
#   )
# fi

# res=$(
#   exec /bin/goose -dir /app/migrations/ sqlite3 "$DB_PATH" up
# )

if [ -f "/bin/air" ]; then
  echo "air installed"
  # exec /bin/litestream replicate -exec "/bin/air"
  exec "/bin/air"
else
  echo "no air found"
  # exec /bin/litestream replicate -exec "/app/main"
  exec "/app/main"
fi

