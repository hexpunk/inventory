#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

# Check for migration name argument
if [ -z "$(echo "$@" | xargs)" ]; then
  echo "missing migration name"

  exit 1
fi

# Slugify
  # Transliterate everything to ASCII
  # Strip out apostrophes
  # Anything that's not a letter or number to a dash
  # Strip leading & trailing dashes
  # Everything to lowercase
function slugify() {
  iconv -t ascii//TRANSLIT \
  | tr -d "'" \
  | sed -E 's/[^a-zA-Z0-9]+/-/g' \
  | sed -E 's/^-+|-+$//g' \
  | tr "[:upper:]" "[:lower:]"
}

PROJECT_ROOT=$(realpath "$(dirname "$0")/..")
MIRGRATION_DIR=$PROJECT_ROOT/migrations
SECS_SINCE_EPOCH=$(date +%s)
MIGRATION_FILE=$MIRGRATION_DIR/$SECS_SINCE_EPOCH-$(echo "$@" | slugify).sql

mkdir -p "$MIRGRATION_DIR" || exit 1

touch "$MIGRATION_FILE" || exit 1

echo + "$MIGRATION_FILE"
