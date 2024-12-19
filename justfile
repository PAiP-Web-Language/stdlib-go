set shell := ["nu", "-c"]
mod? env

default:
  just --list --list-submodules

# Run testing code (in main.go)
run:
    go run ./main.go

# Test
test-watch:
    # go test -v -count=1 "./..."
    go install gotest.tools/gotestsum@latest
    gotestsum -f testname -- -v -count=1 "./..."
    gotestsum -f testname --watch  -- -v -count=1 "./..."

test:
    go install gotest.tools/gotestsum@latest
    gotestsum -f testname -- -v -count=1 "./..."
