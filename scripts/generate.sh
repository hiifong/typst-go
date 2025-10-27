#!/usr/bin/env bash
set -euxo pipefail
THIS_SCRIPT_DIR="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

RUST_FOLDER="$THIS_SCRIPT_DIR/../typst-rs"
GO_FOLDER="$THIS_SCRIPT_DIR/../"

TARGETS="x86_64-unknown-linux-gnu aarch64-unknown-linux-gnu aarch64-apple-darwin x86_64-apple-darwin"

if ! command -v cross &> /dev/null; then
    echo "▸ Cross tool not found, installing"
    cargo install cross --git https://github.com/cross-rs/cross
fi

if ! command -v rustup &> /dev/null; then
    echo "▸ Rustup not found, please install it"
    exit 1
else
	rustup target add $TARGETS
fi

echo "▸ Generate Go bindings"
cd "$RUST_FOLDER"
cargo run \
    --features=cli \
    --bin uniffi-bindgen-go \
    "$RUST_FOLDER/src/typst.udl" \
    --out-dir "target/go"

cp -r "${RUST_FOLDER}/target/go/typst/" "${GO_FOLDER}"

export RUSTFLAGS="-C target-feature=+crt-static"

for TARGET in $TARGETS; do
  echo "▸ Building for $TARGET"
  LIB_NAME="libtypst.so"
  if [[ $TARGET == *-apple-darwin ]]; then
    LIB_NAME="libtypst.dylib"
  fi

	cross build --manifest-path "./Cargo.toml" --target "$TARGET" --locked --release

	mkdir -p "${GO_FOLDER}/libs/${TARGET}"
	cp "${RUST_FOLDER}/target/${TARGET}/release/${LIB_NAME}" "${GO_FOLDER}/libs/${TARGET}/${LIB_NAME}"
done
