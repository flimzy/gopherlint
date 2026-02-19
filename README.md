# `gopherlint`

A stable, bugfix-only fork of [`golangci-lint`](https://github.com/golangci/golangci-lint).

## Purpose

`gopherlint` exists for teams that want a linter that doesn't change between releases. `golangci-lint`
is an excellent project that moves quickly — adding linters, improving defaults, and keeping pace
with the Go ecosystem. This fork trades that freshness for predictability.

**What changes:** Only backported bugfixes from upstream, or original fixes for regressions. No
new features, no linter additions or removals, no config-breaking changes.

**What doesn't change:** The linter set, configuration format, behavior, and CLI interface remain
frozen for the life of a major version.

## Based on

`golangci-lint` v2.10.1

## Versioning

```
v1.GOMAJOR.PATCH
```

| Segment    | Meaning                                                    |
|------------|------------------------------------------------------------|
| `1`        | Fork major version — stable, no breaking changes           |
| `GOMAJOR`  | Major Go version this release targets (e.g. `26` = Go 1.26) |
| `PATCH`    | Bugfix increment                                           |

Example: `v1.26.0` is the first release targeting Go 1.26, based on `golangci-lint` v2.10.1.

## Installation

Binary releases are available on the [Releases page](https://github.com/flimzy/gopherlint/releases).

> **Note:** `go install` is not supported yet. The module path has not been updated from the
> upstream (`github.com/golangci/golangci-lint/v2`), so `go install
> github.com/flimzy/gopherlint@...` will not work. Use the binary releases.

Download the archive for your platform, extract it, and place the `gopherlint` binary somewhere on
your `PATH`.

## Usage

`gopherlint` is a drop-in replacement for `golangci-lint`. It accepts the same flags and the same
configuration format. The only difference is the binary name. Refer to the
[`golangci-lint` documentation](https://golangci-lint.run/docs/) — it applies directly to
`gopherlint` (based on v2.10.1).

```sh
# Same config, same flags — just a different binary name
gopherlint run ./...
```

Existing `.golangci.yml` / `.golangci.toml` / `.golangci.json` config files work without
modification.

## GitHub Actions

The [official `golangci-lint` action](https://github.com/golangci/golangci-lint-action) hardcodes
the binary name `golangci-lint`, so it cannot invoke `gopherlint` directly. Two approaches are
available.

### Option A: run directly (simpler)

```yaml
- name: Install gopherlint
  run: |
    curl -sSfL https://github.com/flimzy/gopherlint/releases/download/v1.26.0/gopherlint-v1.26.0-linux-amd64.tar.gz \
      | tar -xz --strip-components=1 -C /usr/local/bin gopherlint-v1.26.0-linux-amd64/gopherlint

- name: Run gopherlint
  run: gopherlint run ./...
```

### Option B: use the official action (GitHub annotations)

Install `gopherlint` and expose it as `golangci-lint` so the official action can find and invoke
it. This preserves inline PR annotations.

```yaml
- name: Install gopherlint
  run: |
    curl -sSfL https://github.com/flimzy/gopherlint/releases/download/v1.26.0/gopherlint-v1.26.0-linux-amd64.tar.gz \
      | tar -xz --strip-components=1 -C /usr/local/bin gopherlint-v1.26.0-linux-amd64/gopherlint
    ln -s /usr/local/bin/gopherlint /usr/local/bin/golangci-lint

- name: Run gopherlint via golangci-lint action
  uses: golangci/golangci-lint-action@v6
  with:
    install-mode: none
```

## New features and linters

This fork does not add new features or linters. For that, use
[`golangci-lint`](https://github.com/golangci/golangci-lint) directly.

## License

[GPL-3.0](LICENSE)
