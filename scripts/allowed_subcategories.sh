#!/bin/bash

# This script makes sure that we do not accidentally introduce a new subcategory in the docs
# Each subcategory is rendered as a separate list in the docs page at https://registry.terraform.io/providers/e-breuninger/netbox/latest/docs
# The list in the docs should be concise, not cluttered with typos or other inadvertent subcategories
#
# The list of allowed subcategories is maintained in a file (see below)

readonly allowfilepath=".github/allowed-subcategories.txt"

while read -r line; do
    if ! grep --quiet "$line" "$allowfilepath"; then
        echo "error: subcategory \"$line\" is not in $allowfilepath"
        exit 1
    fi
done <<<"$(grep --no-filename --recursive subcategory docs | sort --unique | cut --delimiter='"' --fields=2)"
