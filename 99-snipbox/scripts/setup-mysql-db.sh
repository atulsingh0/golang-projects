#!/bin/sh

if command -v "podman" >/dev/null 2>&1; then
  TOOL="podman"
else
  TOOL="docker"
fi

$TOOL run -d -it --env MYSQL_ROOT_PASSWORD="Passw0rd1" --name mysqlDB -p 3306:3306 -d mysql:latest
echo Sleeping for 10 sec && sleep 10

$TOOL exec -i mysqlDB mysql -uroot -pPassw0rd1 <"${PWD}/sql-commands.sql"
