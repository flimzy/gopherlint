# Changelog

## v1.27.1 (unreleased)

### Bug Fixes (backported)

- **unparam**: Fix panic on type parameters used by closures ([mvdan/unparam#2fa3d84](https://github.com/mvdan/unparam/commit/2fa3d841b0c8)). Packages that previously crashed now lint normally, so new findings may appear in generic code. Parameters of zero-sized generic types with phantom type parameters are no longer reported, matching `unparam`'s zero-size skip rule.

## v1.27.0 (2026-08-20)

Rebased onto `golangci-lint` v2.13.1, targeting Go 1.27. This pulls forward everything from
v2.11.0 through v2.13.1, including go1.27 support, new linters (`clickhouselint`, `exhaustruct_v5`,
`gomodguard_v2`), and the switch to an embedded JSON schema for `config verify`.

### Notes

- **config verify**: Upstream now embeds the JSON schema in the binary instead of fetching it by
  URL, which made gopherlint's custom schema-URL patch (from v1.26.0) obsolete; it has been dropped.
- **sqlclosecheck**: The vendored fork was rebased onto upstream v0.6.0 (new `NewDeferOnlyAnalyzer`
  API) with gopherlint's two local fixes reapplied: the nil-`Referrers()` guard and the pgx
  interface-method-dispatch fix for `Close()` detection.

## v1.26.3 (2026-05-04)

### Bug Fixes (backported)

- **gopherlint custom**: Filter env vars when cloning the repository in custom build ([golangci-lint#6515](https://github.com/golangci/golangci-lint/pull/6515))
- **godot**: Correct auto-fix replacement offset for inline comments ([v1.5.5](https://github.com/tetafro/godot/releases/tag/v1.5.5))
- **godot**: Prevent panic when slicing source lines on files with `//line` directives ([v1.5.6](https://github.com/tetafro/godot/releases/tag/v1.5.6))
- **sqlclosecheck**: Recognize `Close()` called via interface method dispatch on pgx target types ([golangci/sqlclosecheck#1b5fadb](https://github.com/golangci/sqlclosecheck/commit/1b5fadbb85650600cf4e972692888780f5344148))
- **gosec**: Fix G115 false positives for guarded int64-to-byte conversions ([securego/gosec#1578](https://github.com/securego/gosec/pull/1578))
- **gosec**: Per-package rule instantiation eliminates concurrent map crash in G304 ([securego/gosec#1589](https://github.com/securego/gosec/pull/1589))
- **gosec**: Skip SSA analysis on ill-typed packages to prevent panic ([securego/gosec#1607](https://github.com/securego/gosec/pull/1607))

## v1.26.2 (2026-03-23)

### Bug Fixes (backported)

- **fatih/color**: Fix nil os.Stdout panic on Windows services, fix escape code byte counts in Fprint/Fprintf ([v1.19.0](https://github.com/fatih/color/releases/tag/v1.19.0))
- **sqlclosecheck**: Fix nil pointer panic when Referrers() returns nil ([golangci-lint#6439](https://github.com/golangci/golangci-lint/issues/6439))

## v1.26.1 (2026-03-12)

### Bug Fixes (backported)

- **gosec**: Fix panic on float constants in overflow analyzer ([securego/gosec#1505](https://github.com/securego/gosec/pull/1505))
- **gosec**: Fix panic when scanning multi-module repos from root ([securego/gosec#1504](https://github.com/securego/gosec/pull/1504))
- **gosec**: Fix G407 incorrect detection of fixed IV ([securego/gosec#1509](https://github.com/securego/gosec/pull/1509))
- **gosec**: Fix G115 false positives and false negatives ([securego/gosec#1518](https://github.com/securego/gosec/pull/1518))
- **gosec**: Fix G602 false positive for array element access ([securego/gosec#1499](https://github.com/securego/gosec/pull/1499))
- **gosec**: Fix G704 false positive on const URLs ([securego/gosec#1551](https://github.com/securego/gosec/pull/1551))
- **gosec**: Fix Sonar report schema compliance ([securego/gosec#1507](https://github.com/securego/gosec/pull/1507))
- **gosec**: Fix SARIF output invalid null relationships ([securego/gosec#1569](https://github.com/securego/gosec/pull/1569))
- **unqueryvet**: Fix false positives on constants, prepared statements, and broad name patterns ([v1.5.4](https://github.com/MirrexOne/unqueryvet/releases/tag/v1.5.4))

## v1.26.0 (2026-02-19)

Initial release. Based on `golangci-lint` v2.10.1, targeting Go 1.26.
