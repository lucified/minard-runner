#!/bin/bash

set -e

conf=/etc/gitlab-runner/config.toml

if [ "$1" = 'gitlab-runner'  ] && [ ! -f "$conf" ]; then
  host="${PGHOST:-postgresql}"
  user="${PGUSER:-gitlab}"
  password="${PGPASSWORD:-password}"
  db="${PGDATABASE:-gitlabhq_production}"

  export PGHOST="$host"
  export PGUSER="$user"
  export PGPASSWORD="$password"
  export PGDATABASE="$db"

  until psql -q -c '\l' &> /dev/null; do
    echo "Postgres is unavailable - sleeping"
    sleep 2
  done

  >&2 echo "Postgres is up - fetch runner registration token"

  token=`psql -q -A -t -c 'select runners_registration_token from application_settings;'`
  until gitlab-runner register -n -r $token; do
    echo "Gitlab is unavailable - sleeping"
    sleep 2
  done
fi

exec "$@"

