#!/usr/bin/env bash
# Run the legacy and modern error-handling programs from
# test/errorhandling.dir for GOOS/GOARCH, and check that both print the same
# trace as the native legacy program. Binaries for other architectures run
# through binfmt (for example QEMU); wasm uses the lib/wasm exec wrappers.
#
# usage: misc/gon/cross_pair.sh GOOS GOARCH
set -euo pipefail

goos=$1
goarch=$2
root=$(cd "$(dirname "$0")/../.." && pwd)
go="$root/bin/go"
dir="$root/test/errorhandling.dir"
unset GOROOT GOFLAGS
export GOTOOLCHAIN=local PATH="$root/lib/wasm:$PATH"

want=$("$go" run "$dir/common.go" "$dir/legacy.go")

check() {
	local name=$1
	shift
	local got
	got=$(GOOS=$goos GOARCH=$goarch "$go" run "$@")
	if [ "$got" != "$want" ]; then
		echo "$name on $goos/$goarch differs from the native legacy program:"
		diff <(echo "$want") <(echo "$got") || true
		exit 1
	fi
	echo "ok: $name on $goos/$goarch"
}

check legacy "$dir/common.go" "$dir/legacy.go"
check modern "$dir/common.go" "$dir/modern.go"
check "modern without inlining" -gcflags=-l "$dir/common.go" "$dir/modern.go"
