#!/bin/bash

branch_name=$(git symbolic-ref --short HEAD 2>/dev/null || git rev-parse --short HEAD 2>/dev/null)
if [ -z "$branch_name" ]; then
  exit 0
fi

commit_msg_file="$1"
current_msg=$(cat "$commit_msg_file")

if [[ "$current_msg" != "[$branch_name]: "* ]]; then
  echo "[$branch_name]: $current_msg" > "$commit_msg_file"
fi