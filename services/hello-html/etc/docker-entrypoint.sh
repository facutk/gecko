#!/usr/bin/with-contenv bash

set -e # Exit build script on first failure.
set -u # Exit on unset variable.
set -x # Echo commands to stdout.

ls -la
# chmod -R 777 ./

echo "hello-html: installing, this will take a while..."
npm i --loglevel verbose
npm start
