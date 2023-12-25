#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

# Check if an argument is provided
if [ $# -eq 0 ]; then
    echo "Error: No argument provided. Please provide a file." >&2
    exit 1
fi

# Check if the provided argument is a file
file_path=$1
if [ ! -f "$file_path" ]; then
    echo "Error: '$file_path' is not a valid file." >&2
    exit 1
fi

function randint() {
  echo $(( RANDOM % ($2 - $1 + 1) + $1 ))
}

function randwords() {
  if [ "$1" -eq 0 ]; then
    return
  fi

  grep -E -v "'s$" /usr/share/dict/words | shuf -n "$1" | tr -s '\n' ' ' | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//'
}

function insert_item() {
  local name
  name=$(randwords "$(randint 2 4)")

  local description
  description=$(randwords "$(randint 0 30)")

  local quantity
  quantity=$(randint 1 5)

  sqlite3 "$1" "INSERT INTO items (name, description, quantity) VALUES (\"$name\", \"$description\", $quantity);"
}

function insert_container() {
  local name
  name=$(randwords "$(randint 2 4)")

  local description
  description=$(randwords "$(randint 0 30)")

  sqlite3 "$1" "INSERT INTO containers (name, description) VALUES (\"$name\", \"$description\");"
}

for _ in $(seq "$(randint 20 50)"); do
  insert_item "$file_path"
done

for _ in $(seq "$(randint 20 50)"); do
  insert_container "$file_path"
done
