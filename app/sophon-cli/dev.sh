#!/usr/bin/env bash

OUT="${PLANDEX_DEV_CLI_OUT_DIR:-/usr/local/bin}"
NAME="${PLANDEX_DEV_CLI_NAME:-sophon-dev}"
ALIAS="${PLANDEX_DEV_CLI_ALIAS:-sdxd}"

# Double quote to prevent globbing and word splitting.
go build -o "$NAME" &&
    rm -f "$OUT"/"$NAME" &&
    cp "$NAME" "$OUT"/"$NAME" &&
    ln -sf "$OUT"/"$NAME" "$OUT"/"$ALIAS" &&
    echo built "$NAME" cli and added "$ALIAS" alias to "$OUT"
