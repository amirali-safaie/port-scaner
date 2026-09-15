# port-scaner

A concurrent TCP port scanner written in Go. Supports scanning single IPs or entire subnets (CIDR notation) with configurable worker pools.

## Features

- **Single IP or CIDR scanning** - Scan a single host or an entire subnet
- **Concurrent workers** - Configurable parallel scanning with worker pools
- **TCP connect scan** - Uses TCP connect to detect open ports
- **Clean output** - Simple terminal display of results

## Project Structure

```
.
├── main.go                  # Entry point
├── config.json              # Worker pool configuration
├── go.mod                   # Go module definition
├── Makefile                 # Build and test commands
├── extractor/
│   ├── extractor.go         # IP/CIDR host extraction
│   └── extract_host_test.go
├── scaner/
│   ├── send_pack.go         # Scan orchestration
│   ├── worker.go            # Worker pool and TCP connection
│   └── worker_test.go
├── presenter/
│   └── present.go           # Output formatting
└── utiles/
    ├── config.go            # Config file reader
    ├── utiles.go            # IP math utilities
    └── utiles_test.go
```

## Getting Started

### Prerequisites

- Go 1.22 or later

### Build

```bash
make build
```

### Run

```bash
make run
```

### Test

```bash
make test
```

### Other Commands

```bash
make check    # Format, vet, and test
make fmt      # Format code
make vet      # Run go vet
make clean    # Remove build artifacts
```

## Usage

When you run the scanner, it will prompt you for:

1. **Target** - An IP address (`192.168.1.1`) or CIDR range (`192.168.1.0/24`)
2. **Port range** - A range in `start-end` format (e.g., `1-1024`)

Example:

```
Enter the target: 192.168.1.0/30
Determine the port range (start-end): 80-443
```

## Configuration

Edit `config.json` to adjust the number of concurrent workers:

```json
{
  "worker_num": 20
}
```

## Architecture

The scanner uses a pipeline architecture:

1. **Extractor** - Resolves the target (single IP or CIDR) into individual IPs
2. **Scanner** - Dispatches scan jobs to a pool of workers that attempt TCP connections
3. **Presenter** - Displays results as they come in

Channels connect the stages:

```
Extractor → IPs channel → Scanner → Results channel → Presenter
```

## License

This project is for learning purposes.
