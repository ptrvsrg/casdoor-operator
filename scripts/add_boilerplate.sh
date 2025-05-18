#!/bin/bash

BOILERPLATE_FILE="hack/boilerplate.go.txt"

if [[ ! -f "$BOILERPLATE_FILE" ]]; then
    echo "File $BOILERPLATE_FILE not found!"
    exit 1
fi

BOILERPLATE_LINES=$(wc -l < "$BOILERPLATE_FILE")

find . -type f -name "*.go" | while read -r file; do
    if grep -q "DO NOT EDIT" "$file"; then
        echo "File $file generated (found DO NOT EDIT), miss"
        continue
    fi

    if ! diff -q <(head -n "$BOILERPLATE_LINES" "$file" 2>/dev/null) <(head -n "$BOILERPLATE_LINES" "$BOILERPLATE_FILE" 2>/dev/null) >/dev/null 2>&1; then
        echo "Added boilerplate в $file"
        {
            cat "$BOILERPLATE_FILE"
            printf "\n\n"
            cat "$file"
        } > "${file}.tmp" && mv "${file}.tmp" "$file"
    else
        echo "Boilerplate already exists in $file"
    fi
done