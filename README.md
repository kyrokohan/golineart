## GoLineArt (GLA)

#### Turn images into grayscale line-art by iteratively drawing the best line that reduces error versus the target image.

![](./examples/sistine-chapel.gif)


### Examples

Original                   | Output 
:-------------------------:|:-------------------------:
![](./examples/originals/mona-lisa.jpg)  |  ![](./examples/outputs/mona-lisa.jpeg)

> More examples are available under the `examples` directory. These examples were generated using the following configuration:
> `-alpha 10 -rounds 100000 -lines 100`

### Install

#### Option 1: Download pre-built binaries
Download the latest release from the [Releases page](https://github.com/kyrokohan/golineart/releases).

```bash
# Linux/macOS: Make the binary executable
chmod +x golineart-<os>-<arch>

# Example for macOS on Apple Silicon
chmod +x golineart-darwin-arm64
```

#### Option 2: Build from source
- Requires a recent Go toolchain (module sets `go 1.24.4`).
- Build the CLI:
```bash
./scripts/build.sh
# or
go build -o bin/gla cmd/gla/main.go
```

### Quick start

#### Using pre-built binaries
```bash
# macOS (Apple Silicon)
./golineart-darwin-arm64 path/to/image.jpg

# Linux (x86_64)
./golineart-linux-amd64 path/to/image.jpg
```

```powershell
# Windows (PowerShell)
.\golineart-windows-amd64.exe path\to\image.jpg
```

#### Using locally built binary
```bash
./bin/gla path/to/image.jpg
```

Defaults: `-rounds 100000`, `-lines 100`, `-alpha 10`, `-odir out`, `-ofile final`, `-oext jpeg`.

### Usage
```bash
# For pre-built binaries
golineart-<os>-<arch> [flags] <image>

# For locally built binary
./bin/gla [flags] <image>
```
- `-rounds int`  total rounds (lines) to draw (default 100000)
- `-lines int`   candidates per round (default 100)
- `-alpha uint`  opacity [0–255] (default 10)
- `-sfreq int`   save every N rounds (0 disables)
- `-odir string` output directory (default `out`)
- `-ofile string` output base name (default `final`)
- `-oext string` output format: `png`, `jpg`, or `jpeg` (default `jpeg`)

Examples:
```bash
# Higher detail, save progress frames every 200 rounds, PNG output
./bin/gla -rounds 30000 -lines 500 -sfreq 200 -oext png path/to/image.png

# Different alpha
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

## Releases

See the [Releases page](https://github.com/kyrokohan/golineart/releases) for download links and release notes.

## License

This project is licensed. See `LICENSE` for details.