#!/bin/bash

. /usr/local/bin/docker-entrypoint.sh

docker_setup_env >/dev/null
docker_create_db_directories >/dev/null

if [ "$(id -u)" = '0' ]; then
  exec gosu postgres "$BASH_SOURCE" "$@"
fi

docker_init_database_dir >/dev/null 2>&1
pg_setup_hba_conf >/dev/null 2>&1
postgres >/dev/null 2>&1 &

while ! pg_isready >/dev/null; do
  sleep 0.1
done

cd "/code/$1" || exit

psql -q -f schema.sql

time psql -q -t -f run.sql
