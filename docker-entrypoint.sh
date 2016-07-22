#!/bin/bash

#set -e

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

  token=""
  min_length=15

  until [ "${#token}" -gt "$min_length" ] && gitlab-runner register -n -r $token >/dev/null 2>&1; do
    echo "Failed to register - sleeping"
    sleep 5
    token=`psql -q -A -t -c 'select runners_registration_token from application_settings;' 2>/dev/null`
  done

  echo "Runner registered succesfully"

fi

exec "$@"

