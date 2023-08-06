#!/bin/sh
cur_path="$(cd "$(dirname "$0")" && pwd)"

if command -v "podman" >/dev/null 2>&1; then
  TOOL="podman"
else
  TOOL="docker"
fi

$TOOL rm mysqlDB -f
