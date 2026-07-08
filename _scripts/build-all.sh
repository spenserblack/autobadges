#!/bin/sh
SCRIPT="$(readlink -f "$0")"
BASE_DIR="$(dirname "$(dirname "$SCRIPT")")"
cd "$BASE_DIR"
DIST_DIR="$BASE_DIR/dist"
echo "Building to $DIST_DIR"
PLATFORMS="darwin-arm64 linux-amd64 windows-amd64"
for platform in $PLATFORMS; do
	echo "Building for $platform"
	goos="$(echo "$platform" | cut -d - -f 1)"
	goarch="$(echo "$platform" | cut -d - -f 2)"
	exe_name="autobadges"
	if [ "$goos" = "windows" ]; then
		exe_name="$exe_name.exe"
	fi
	target="$DIST_DIR/$exe_name"
	GOOS="$goos" GOARCH="$goarch" go build -o "$target" ./autobadges.go

	tar -C "$DIST_DIR" --remove-files -czf "$DIST_DIR/autobadges-$platform.tar.gz" "$exe_name"
done
