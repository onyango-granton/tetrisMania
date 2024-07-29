#!/bin/bash

# Define the directories
directories=("gemini" "gtTest" "chatgpt" "claude" "claude2" "gtTest")

# Iterate over each directory
for dir in "${directories[@]}"; do
    # Check if directory exists
    if [ -d "$dir" ]; then
        # Iterate over each file in the directory
        for file in "$dir"/*; do
            # Check if it's a file (not a directory)
            if [ -f "$file" ]; then
                # Add the file to git
                git add -f "$file"
                # Commit the file with a message
                git commit -m "feat: add $(basename "$file")"
            fi
        done
    else
        echo "Directory $dir does not exist"
    fi
done
