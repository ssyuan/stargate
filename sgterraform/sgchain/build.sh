#!/bin/sh
set -e

OUT_DIR="${1?[Specify relative output directory]}"
shift
LIBRA_DIR="$(cd ./libra && pwd)"
TF_WORK_DIR="$(cd ./sgterraform/sgchain && pwd)"
OUTPUT_DIR="$TF_WORK_DIR/$OUT_DIR"
mkdir -p "$OUTPUT_DIR"


if [ ! -e "$OUTPUT_DIR/mint.key" ]; then
	cd $LIBRA_DIR && cargo run -p generate-keypair --bin generate-keypair -- -o "$OUTPUT_DIR/mint.key"
fi

cd $LIBRA_DIR && cargo run -p config-builder --bin libra-config -- -b $LIBRA_DIR/config/data/configs/node.config.toml -m "$OUTPUT_DIR/mint.key" -o "$OUTPUT_DIR/val" -d -r validator "$@"


