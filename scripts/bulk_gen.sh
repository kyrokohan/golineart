#!/usr/bin/env bash
set -Eeuo pipefail

usage() {
  cat >&2 <<'EOF'
Usage:
  run_gla.sh <folder> [options]

Options (forwarded to ./bin/gla if provided):
  -alpha N   | --alpha N
  -lines N   | --lines N
  -rounds N  | --rounds N
  -oext EXT  | --oext EXT
  -sfreq N   | --sfreq N
  -odir DIR  | --odir DIR

Notes:
- Only .jpeg and .png files are processed (recursively).
- Each image's base filename (without extension) is passed as -ofile.
EOF
}

dir=""
alpha=""; lines=""; rounds=""; oext=""; sfreq=""; odir=""

# Parse args (accepts -flag and --flag forms)
while [[ $# -gt 0 ]]; do
  case "$1" in
    -h|--help) usage; exit 0 ;;
    -alpha|--alpha)   alpha="${2:?missing value for $1}"; shift 2 ;;
    -lines|--lines)   lines="${2:?missing value for $1}"; shift 2 ;;
    -rounds|--rounds) rounds="${2:?missing value for $1}"; shift 2 ;;
    -oext|--oext)     oext="${2:?missing value for $1}"; shift 2 ;;
    -sfreq|--sfreq)   sfreq="${2:?missing value for $1}"; shift 2 ;;
    -odir|--odir)     odir="${2:?missing value for $1}"; shift 2 ;;
    --) shift; break ;;
    -*)
      echo "Unknown option: $1" >&2
      usage; exit 1 ;;
    *)
      if [[ -z "$dir" ]]; then dir="$1"; else
        echo "Unexpected extra argument: $1" >&2
        usage; exit 1
      fi
      shift ;;
  esac
done

if [[ -z "$dir" ]]; then
  echo "Error: folder argument is required." >&2
  usage; exit 1
fi
if [[ ! -d "$dir" ]]; then
  echo "Error: '$dir' is not a directory." >&2
  exit 1
fi
if [[ ! -x ./bin/gla ]]; then
  echo "Error: ./bin/gla not found or not executable from $(pwd)." >&2
  exit 1
fi

# If an output directory was provided, ensure it exists
if [[ -n "$odir" ]]; then
  mkdir -p "$odir"
fi

# Build args to forward to ./bin/gla
gla_args=()
[[ -n "$alpha"  ]] && gla_args+=(-alpha  "$alpha")
[[ -n "$lines"  ]] && gla_args+=(-lines  "$lines")
[[ -n "$rounds" ]] && gla_args+=(-rounds "$rounds")
[[ -n "$oext"   ]] && gla_args+=(-oext   "$oext")
[[ -n "$sfreq"  ]] && gla_args+=(-sfreq  "$sfreq")
[[ -n "$odir"   ]] && gla_args+=(-odir   "$odir")

# Find only .jpeg and .png (case-insensitive), recursively, safely handle spaces/NULs
find "$dir" -type f \( -iname '*.jpeg' -o -iname '*.png' -o -iname '*.jpg' \) -print0 |
while IFS= read -r -d '' img; do
  base="$(basename "$img")"
  name_no_ext="${base%.*}"   # e.g., "photo" from "photo.png"

  echo "Processing: $img"
  ./bin/gla "${gla_args[@]}" -ofile "$name_no_ext" "$img"
done
