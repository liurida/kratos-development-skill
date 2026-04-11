#!/bin/bash
# This script links all Kratos skills from this repository into a target project's
# .claude/skills directory using absolute symbolic links.

set -e

if [ -z "$1" ]; then
    echo "Usage: $0 <ABSOLUTE_path_to_target_project>"
    echo "Example: $0 /path/to/your/project"
    exit 1
fi

# --- Path Validation ---
# Check if the provided path is absolute. This is required for robustness.
if [[ "$1" != /* ]]; then
    echo "Error: The path you provided is not an absolute path."
    echo "Please provide the full, absolute path to your target project."
    echo "Example: $0 /path/to/your/project"
    exit 1
fi

# Get the absolute path of the directory containing this script.
SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)

# --- Target Directory Logic ---
# Check if the user provided the path to the project or directly to the skills directory.
if [[ "$1" == *"/.claude/skills" ]]; then
    TARGET_SKILLS_DIR="$1"
else
    TARGET_SKILLS_DIR="$1/.claude/skills"
fi

# Ensure the final target directory exists.
if ! [ -d "$(dirname "$TARGET_SKILLS_DIR")" ]; then
    echo "Error: The parent directory for the target skills folder does not exist."
    echo "Please check the path: $1"
    exit 1
fi
mkdir -p "$TARGET_SKILLS_DIR"

echo "Linking Kratos skills into $TARGET_SKILLS_DIR..."

# Find all skill directories (kratos-*) in the source location.
for skill_path in "$SCRIPT_DIR"/kratos-*; do
    if [ -d "$skill_path" ]; then
        skill_name=$(basename "$skill_path")
        target_link_path="$TARGET_SKILLS_DIR/$skill_name"

        if [ -L "$target_link_path" ]; then
            echo "Link for '$skill_name' already exists, skipping."
        else
            echo "Linking '$skill_name'..."
            # Use the absolute path for the source of the link
            ln -s "$skill_path" "$target_link_path"
        fi
    fi
done

echo "Done. All Kratos skills are now available in your project."
