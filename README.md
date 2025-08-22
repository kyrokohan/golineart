## GoLineArt (GLA)

Turn images into grayscale line-art by iteratively drawing the best line that reduces error versus the target image.

### Install
- Requires a recent Go toolchain (module sets `go 1.24.4`).
- Build the CLI:
```bash
./scripts/build.sh
# or
go build -o bin/gla cmd/gla/main.go
```

### Quick start
```bash
./bin/gla path/to/image.jpg
```
Defaults: `-rounds 25000`, `-lines 200`, `-alpha 51`, `-odir out`, `-ofile final`, `-oext jpeg`.

### Usage
```bash
./bin/gla [flags] <image>
```
- `-rounds int`  total rounds (lines) to draw (default 25000)
- `-lines int`   candidates per round (default 200)
- `-alpha uint`  opacity [0–255] (default 51)
- `-sfreq int`   save every N rounds (0 disables)
- `-odir string` output directory (default `out`)
- `-ofile string` output base name (default `final`)
- `-oext string` output format: `png`, `jpg`, or `jpeg` (default `jpeg`)

Examples:
```bash
# Higher detail, save progress frames every 200 rounds, PNG output
./bin/gla -rounds 30000 -lines 500 -sfreq 200 -oext png path/to/image.png

# Subtler lines
./bin/gla -alpha 32 path/to/image.jpg
```

### Batch process a folder
Process all `.jpg`, `.jpeg`, and `.png` files recursively:
```bash
./scripts/bulk_gen.sh <folder> [GLA flags...]
```
All args after `<folder>` are forwarded directly to `./bin/gla`.

Examples:
```bash
./scripts/bulk_gen.sh ./photos --rounds 20000 --lines 400 --alpha 64 --oext png --odir out
./scripts/bulk_gen.sh ./photos -rounds 15000 -sfreq 250
```

### How it works (1‑minute tour)
- Converts the input to grayscale and creates a white canvas of the same size.
- For each round, samples N random edge-to-edge lines, picks the one that best lowers MSE, and draws it with the given opacity.
- Uses all CPU cores to evaluate candidates per round.

## License

This project is licensed. See `LICENSE` for details.