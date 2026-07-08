#!/bin/sh
INSTALL_DIR="/usr/local/bin"
os="$(uname -s)"
arch="$(uname -m)"

case "$os" in
	Linux)
		goos="linux"
		;;
	Darwin)
		goos="darwin"
		;;
	*)
		echo "Unsupported OS: $os" >&2
		exit 1
		;;
esac
case "$arch" in
	x86_64)
		goarch="amd64"
		;;
	amd64)
		goarch="amd64"
		;;
	arm64)
		goarch="arm64"
		;;
	*)
		echo "Unsupported machine: $arch" >&2
		exit 1
		;;
esac
asset_name="autobadges-$goos-$goarch"

extract_target="$(mktemp -d)"
echo "Downloading and unpacking to '$INSTALL_DIR'"
curl -fsSL "https://github.com/spenserblack/autobadges/releases/latest/download/$asset_name.tar.gz" | tar -C "$INSTALL_DIR" -xzf -
