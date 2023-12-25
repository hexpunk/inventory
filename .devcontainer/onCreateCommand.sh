#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

# Install sqlite3 and a /usr/share/dict/words file
sudo apt update
sudo apt install -y sqlite3 wamerican

# Install Bun
curl -fsSL https://bun.sh/install | bash
