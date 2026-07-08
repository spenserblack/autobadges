#!/bin/sh
SCRIPT="$(readlink -f "$0")"
BASE_DIR="$(dirname "$(dirname "$SCRIPT")")"
cd "$BASE_DIR"
if ! command -v gh > /dev/null 2>&1; then
	echo "GitHub CLI (gh) is required" >&2
	exit 1
fi

# NOTE: Ensure there are no unwanted files in dist/
echo "Cleaning $BASE_DIR/dist/"
git clean -dfx dist/

echo "Running build script..."
./_scripts/build-all.sh

echo "Creating draft release..."
gh release create --draft --generate-notes "v0.0.0" "$DIST_DIR/*.tar.gz"
