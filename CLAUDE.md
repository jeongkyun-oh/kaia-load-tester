# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

kaia-load-tester is a blockchain load testing tool for the Kaia network (formerly Klaytn). It uses Boomer (Go port of Locust) to generate concurrent transaction loads against Kaia nodes.

## Build and Development Commands

```bash
# Build the main executable
make build                    # Creates ./build/bin/klayslave with embedded version

# Clean and rebuild
make clean && make build

# Run tests
go test ./...                 # Run all tests
go test -v ./testcase/...     # Run test case tests with verbose output
go test -run TestSpecific     # Run specific test

# Dependency management
go mod tidy                   # Clean up dependencies
go mod download              # Download dependencies

# Fix IDE issues (if IntelliJ/GoLand shows errors but build works)
rm -rf .idea/workspace.xml   # Remove corrupted workspace
go clean -modcache && go mod download  # Clear and rebuild module cache
```

## Architecture

### Core Structure
- **Entry Point**: `klayslave/main.go` - CLI application using urfave/cli
- **Test Cases**: `testcase/` - Each subdirectory implements specific transaction types
- **Account Management**: `klayslave/account/` - Manages test accounts, nonces, and key distribution
- **Client Pool**: `klayslave/clipool/` - Connection pooling for RPC clients
- **Configuration**: `klayslave/config/` - Runtime configuration and flags

### Test Case System
All test cases implement `ExtendedTask` interface and are registered in `testcase/tclist.go`:
- Basic transfers: `transferTxTC`, `ethLegacyTxTC`, `tokenTransferTxTC`
- DEX operations: `limitOrderTxTC`, `marketOrderTxTC`, `stopOrderTxTC`, `tpslOrderTxTC`
- Session transactions: `sessionTxTC` (create/delete sessions)
- Each test case can be weighted for load distribution

### Running Load Tests

1. Start Locust master:
```bash
python3 -m venv venv
source venv/bin/activate
pip3 install locust==1.2.3
locust -f dist/locustfile.py --master
```

2. Start slave(s):
```bash
./build/bin/klayslave \
    --max-rps 150 \
    --master-host localhost \
    --master-port 5557 \
    -key $PRIVATE_KEY \
    -tc="transferTxTC,ethLegacyTxTC" \
    -endpoint $RPC_ENDPOINT
```

3. Configure load via Locust web UI at http://localhost:8089

## Key Configuration Flags

- `-tc`: Comma-separated test case names to run
- `-endpoint`: Kaia node RPC URL
- `-key`: Rich account private key for funding test accounts
- `--charge`: Amount to fund each test account (default: 100 KAIA)
- `--vusigned`/`--vuunsigned`: Number of signed/unsigned transaction accounts
- `--max-rps`: Request rate limit per slave
- `--chainid`: Chain ID (default: 2018 for Kaia mainnet)

## Important Dependencies

- Uses custom Kaia fork: `github.com/kaiachain/go-ethereum` (see go.mod replace directive)
- Requires Go 1.23.7+
- Python 3 with Locust 1.2.3 for master node

## Account Management Pattern

The system uses hierarchical token distribution:
1. Rich account funds distributor accounts
2. Distributors fund test accounts in parallel
3. Test accounts execute transactions
4. Nonce management: Thread-safe with time-based nonces for DEX operations

## Adding New Test Cases

1. Create new directory under `testcase/`
2. Implement `ExtendedTask` interface with `Init()` and `Run()` methods
3. Register in `testcase/tclist.go` `init()` function
4. Test case will be available via `-tc` flag

## DEX-Specific Features

- Order book operations require proper session management
- Time-based nonces for order uniqueness
- Support for complex order types (TPSL, stop orders)
- Order decoder utility in `order-decoder/` for debugging