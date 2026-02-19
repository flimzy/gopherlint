.DEFAULT_GOAL = test
.PHONY: FORCE

# enable consistent Go 1.12/1.13 GOPROXY behavior.
export GOPROXY = https://proxy.golang.org

BINARY = gopherlint
ifeq ($(OS),Windows_NT)
	BINARY := $(BINARY).exe
endif

# Build

build: $(BINARY)
.PHONY: build

build_race:
	go build -race -o $(BINARY) ./cmd/gopherlint
.PHONY: build_race

clean:
	rm -f $(BINARY)
	rm -f test/path
.PHONY: clean

# Test

test: export GOLANGCI_LINT_INSTALLED = true
test: CGO_ENABLED=1
test: build
	GL_TEST_RUN=1 ./$(BINARY) run -v
	GL_TEST_RUN=1 go test -v -parallel 2 ./...
.PHONY: test

test_race: build_race
	GL_TEST_RUN=1 ./$(BINARY) run -v --timeout=5m
.PHONY: test_race

# ex: T=output.go make test_integration
# the value of `T` is the name of a file from `test/testdata`
test_integration:
	GL_TEST_RUN=1 go test -v ./test -count 1 -run TestSourcesFromTestdata/$T
.PHONY: test_integration

# ex: T=multiple-issues-fix.go make test_integration_fix
# the value of `T` is the name of a file from `test/testdata/fix`
test_integration_fix: build
	GL_TEST_RUN=1 go test -v ./test -count 1 -run TestFix/$T
.PHONY: test_integration_fix

# Migration

clone_config:
	go run ./pkg/commands/internal/migrate/cloner/
.PHONY: clone_config

# Benchmark

# Benchmark with a local version
# LINTER=gosec VERSION=v1.59.0 make bench_local
bench_local: hyperfine
	@:$(call check_defined, LINTER VERSION, 'missing parameter(s)')
	@./scripts/bench/bench_local.sh $(LINTER) $(VERSION)
.PHONY: bench_local

# Benchmark between 2 existing versions
# make bench_version LINTER=gosec VERSION_OLD=v1.58.2 VERSION_NEW=v1.59.0
bench_version: hyperfine
	@:$(call check_defined, LINTER VERSION_OLD VERSION_NEW, 'missing parameter(s)')
	@./scripts/bench/bench_version.sh $(LINTER) $(VERSION_OLD) $(VERSION_NEW)
.PHONY: bench_version

hyperfine:
	@which hyperfine > /dev/null || (echo "Please install hyperfine https://github.com/sharkdp/hyperfine#installation" && exit 1)
.PHONY: hyperfine

# Non-PHONY targets (real files)

$(BINARY): FORCE
	go build -o $@ ./cmd/gopherlint

go.mod: FORCE
	go mod tidy
	go mod verify
go.sum: go.mod

# Functions

# Check that given variables are set and all have non-empty values,
# die with an error otherwise.
#
# Params:
#   1. Variable name(s) to test.
#   2. (optional) Error message to print.
#
# https://stackoverflow.com/a/10858332/8228109
check_defined = \
    $(strip $(foreach 1,$1, \
        $(call __check_defined,$1,$(strip $(value 2)))))
__check_defined = \
    $(if $(value $1),, \
      $(error Undefined $1$(if $2, ($2))))
