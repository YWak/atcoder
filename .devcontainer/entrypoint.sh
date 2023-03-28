#!/bin/bash

set -e

# modify uid/gid
sudo usermod -u $LOCAL_UID -o user
sudo groupmod -g $LOCAL_GID user

exec "$@"
