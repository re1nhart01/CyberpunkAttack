#!/bin/bash

# Define paths
DOCS_FOLDER="./docs"
CMD_DOCS_FOLDER="./cmd/docs"
MAIN_FILE="cmd/main/main.go"

# Function to check and remove a directory
remove_folder() {
    if [ -d "$1" ]; then
        echo "Removing existing folder: $1"
        rm -rf "$1"
    else
        echo "Folder not found: $1"
    fi
}

# Function to move a directory
move_folder() {
    if [ -d "$1" ]; then
        echo "Moving $1 to $2"
        mv "$1" "$2"
    else
        echo "Source folder not found: $1"
        exit 1
    fi
}

# Remove existing docs folder
remove_folder "$DOCS_FOLDER"

# Generate Swagger documentation
echo "Running swag init..."
swag init -g "$MAIN_FILE" -o "$CMD_DOCS_FOLDER"
if [ $? -ne 0 ]; then
    echo "Error: swag init failed."
    exit 1
fi

# Remove existing cmd/docs folder and move new docs
remove_folder "$DOCS_FOLDER"
move_folder "$CMD_DOCS_FOLDER" "$DOCS_FOLDER"

echo "Swagger documentation generated and moved to $DOCS_FOLDER successfully."