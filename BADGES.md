# Badges

This lists the badges that this project supports, and what triggers a badge to be
included.

## Golang

A Go Reference badge is created if a `go.mod` file is found.

## GitHub

### CLI extensions

If the repository's name starts with `gh-` and a `go.mod` file is found, then the
repository is assumed to be a pre-compiled GitHub CLI extension.
A [shields.io][shields-io] badge will be created that tracks GitHub release asset
downloads.

### Workflows

This project will look for YAML files in `.github/workflows/`. For those that are found
and have the following data, a badge is created.

```yaml
# foo.yml
on:
  push:
    # Replace `DEFAULT` with the name of your default branch.
    branch: [ "DEFAULT" ]
```

## JavaScript (Node.js)

A [shields.io][shields-io] badge will be generated for the latest release on [NPM][npm].

## Ruby

A [badge.fury.io][badge-fury] badge will be created if a `*.gemspec` file is found.

Note that, because a `.gemspec` file is a *script,* not a data file, parsing may be
innacurate. Pay especially close attention to the results if you use this tool on a
Ruby project.

## Rust

[shields.io][shields-io] badges for the latest release on [crates.io][crates-io] and
documentation on [docs.rs][docs-rs] will be created if a `Cargo.toml` file is found.

[badge-fury]: https://badge.fury.io/
[crates-io]: https://crates.io/
[docs-rs]: https://docs.rs/
[npm]: https://www.npmjs.com/
[shields-io]: https://shields.io/
