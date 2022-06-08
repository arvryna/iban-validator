# IBan Validator Service
- Service that can validate IBAN

# Testing:
- Make sure server is running before starting test `go run main.go`
- To run all the tests: `go test -v ./...`

# Usage:
```
curl -X POST -H "Content-Type: application/json" \
    -d '{"iban": "12345"}' \
    http://localhost:9090/validators/iban
```