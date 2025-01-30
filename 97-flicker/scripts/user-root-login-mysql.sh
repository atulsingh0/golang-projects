#!/bin/sh

if command -v "podman" >/dev/null 2>&1; then
  TOOL="podman"
else
  TOOL="docker"
fi

$TOOL exec -it mysqlDB mysql -uroot -pPassw0rd1
