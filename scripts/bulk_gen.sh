#!/usr/bin/env bash
set -Eeuo pipefail

usage() {
  cat >&2 <<'EOF'
Usage:
  bulk_gen.sh <folder> [GLA flags...]

All arguments after <folder> are forwarded directly to ./bin/gla, unchanged.

Notes:
- Only .jpeg, .jpg, and .png files are processed (recursively).
- Each image's base filename (without extension) is passed as -ofile.

Examples:
  ./scripts/bulk_gen.sh ./photos --rounds 20000 --lines 400 --alpha 51 --oext png --odir out
  ./scripts/bulk_gen.sh ./photos -rounds 30000 -sfreq 200 -oext jpeg
EOF
}

if [[ $# -lt 1 ]]; then
  echo "Error: folder argument is required." >&2
  usage; exit 1
fi

if [[ "$1" == "-h" || "$1" == "--help" ]]; then
  usage; exit 0
fi

dir="$1"
shift || true
gla_args=("$@")

if [[ ! -d "$dir" ]]; then
  echo "Error: '$dir' is not a directory." >&2
  exit 1
fi
if [[ ! -x ./bin/gla ]]; then
  echo "Error: ./bin/gla not found or not executable from $(pwd)." >&2
  exit 1
fi

# Find only .jpeg and .png (case-insensitive), recursively, safely handle spaces/NULs
find "$dir" -type f \( -iname '*.jpeg' -o -iname '*.png' -o -iname '*.jpg' \) -print0 |
while IFS= read -r -d '' img; do
  base="$(basename "$img")"
  name_no_ext="${base%.*}"   # e.g., "photo" from "photo.png"

  echo "Processing: $img"
  if (( ${#gla_args[@]} > 0 )); then
    ./bin/gla "${gla_args[@]}" -ofile "$name_no_ext" "$img"
  else
    ./bin/gla -ofile "$name_no_ext" "$img"
  fi
done
