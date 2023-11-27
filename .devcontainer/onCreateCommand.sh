#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

# Install sqlite3
sudo apt update
sudo apt install -y sqlite3

# Install Bun
curl -fsSL https://bun.sh/install | bash
