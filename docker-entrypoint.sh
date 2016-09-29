#!/bin/bash

#set -e

conf=/etc/gitlab-runner/config.toml

gettoken () { psql -q -A -t -c 'select runners_registration_token from application_settings;'; }

if [ "$1" = 'gitlab-runner'  ]; then

  if [[ -z "$DEBUG" ]] ; then
    exec 3>&1 &>/dev/null
  else
    exec 3>&1
  fi

  host="${PGHOST:-postgresql}"
  user="${PGUSER:-gitlab}"
  password="${PGPASSWORD:-password}"
  db="${PGDATABASE:-gitlabhq_production}"

  export PGHOST="$host"
  export PGUSER="$user"
  export PGPASSWORD="$password"
  export PGDATABASE="$db"

  min_length=15
  token=$(gettoken)

  until [ "${#token}" -gt "$min_length" ] && gitlab-runner register -n -r $token; do
    echo "Failed to register - sleeping" >&3
    sleep 5
    token=$(gettoken)
  done

  echo "Runner registered succesfully" >&3

fi

exec "$@"

